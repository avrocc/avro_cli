package registry

import (
	"avro_cli/internal/domain"
	"sort"
	"strings"
	"sync"

	"github.com/sahilm/fuzzy"
)

var (
	global     *Registry
	globalOnce sync.Once
)

// Global returns the singleton registry instance.
func Global() *Registry {
	globalOnce.Do(func() {
		global = New()
	})
	return global
}

// Registry is a thread-safe command catalog.
type Registry struct {
	mu       sync.RWMutex
	commands []domain.CommandDescriptor
}

// New creates an empty registry.
func New() *Registry {
	return &Registry{}
}

// Register adds one or more commands to the catalog.
func (r *Registry) Register(cmds ...domain.CommandDescriptor) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.commands = append(r.commands, cmds...)
}

// All returns every registered command.
func (r *Registry) All() []domain.CommandDescriptor {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]domain.CommandDescriptor, len(r.commands))
	copy(out, r.commands)
	return out
}

// Categories returns sorted unique categories.
func (r *Registry) Categories() []domain.Category {
	r.mu.RLock()
	defer r.mu.RUnlock()

	seen := make(map[string]domain.Category)
	for _, c := range r.commands {
		seen[c.Category.Name] = c.Category
	}

	cats := make([]domain.Category, 0, len(seen))
	for _, cat := range seen {
		cats = append(cats, cat)
	}
	sort.Slice(cats, func(i, j int) bool {
		return cats[i].Name < cats[j].Name
	})
	return cats
}

// ByCategory returns commands in a given category, sorted by name.
func (r *Registry) ByCategory(category string) []domain.CommandDescriptor {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var out []domain.CommandDescriptor
	for _, c := range r.commands {
		if c.Category.Name == category {
			out = append(out, c)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].Name < out[j].Name
	})
	return out
}

// Find looks up a command by category and name.
func (r *Registry) Find(category, name string) (domain.CommandDescriptor, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, c := range r.commands {
		if c.Category.Name == category && (c.Name == name || containsAlias(c.Aliases, name)) {
			return c, true
		}
	}
	return domain.CommandDescriptor{}, false
}

// Search returns commands whose full name or description matches the query (case-insensitive).
func (r *Registry) Search(query string) []domain.CommandDescriptor {
	r.mu.RLock()
	defer r.mu.RUnlock()

	q := strings.ToLower(query)
	var out []domain.CommandDescriptor
	for _, c := range r.commands {
		if strings.Contains(strings.ToLower(c.FullName()), q) ||
			strings.Contains(strings.ToLower(c.Description), q) {
			out = append(out, c)
		}
	}
	return out
}

// commandSource implements fuzzy.Source over a slice of CommandDescriptors.
type commandSource []domain.CommandDescriptor

func (s commandSource) String(i int) string { return s[i].FullName() }
func (s commandSource) Len() int            { return len(s) }

// FuzzySearch returns commands ranked by fuzzy match quality.
func (r *Registry) FuzzySearch(query string) []FuzzyResult {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if query == "" {
		out := make([]FuzzyResult, len(r.commands))
		for i, c := range r.commands {
			out[i] = FuzzyResult{Command: c}
		}
		return out
	}

	src := commandSource(r.commands)
	matches := fuzzy.FindFrom(query, src)

	out := make([]FuzzyResult, len(matches))
	for i, m := range matches {
		out[i] = FuzzyResult{
			Command:        r.commands[m.Index],
			MatchedIndexes: m.MatchedIndexes,
			Score:          m.Score,
		}
	}
	return out
}

// FuzzyResult holds a command and its fuzzy match metadata.
type FuzzyResult struct {
	Command        domain.CommandDescriptor
	MatchedIndexes []int
	Score          int
}

func containsAlias(aliases []string, name string) bool {
	for _, a := range aliases {
		if a == name {
			return true
		}
	}
	return false
}
