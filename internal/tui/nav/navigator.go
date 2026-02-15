package nav

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Screen identifies the current TUI screen.
type Screen int

const (
	HomeScreen Screen = iota
	CategoryScreen
	CommandDetailScreen
	SearchScreen
)

// Entry represents a screen on the navigation stack with context.
type Entry struct {
	Screen Screen
	Title  string
	Data   any // screen-specific context (e.g., category name)
}

// Navigator manages a stack of screens for TUI navigation.
type Navigator struct {
	stack []Entry
}

// New creates a navigator with the home screen.
func New() *Navigator {
	return &Navigator{
		stack: []Entry{{Screen: HomeScreen, Title: "Home"}},
	}
}

// NewWithInitial creates a navigator starting on a custom screen.
func NewWithInitial(screen Screen, title string) *Navigator {
	return &Navigator{
		stack: []Entry{{Screen: screen, Title: title}},
	}
}

// Current returns the top entry.
func (n *Navigator) Current() Entry {
	return n.stack[len(n.stack)-1]
}

// Push adds a new screen.
func (n *Navigator) Push(e Entry) {
	n.stack = append(n.stack, e)
}

// Pop removes the top screen. Returns false if already at root.
func (n *Navigator) Pop() bool {
	if len(n.stack) <= 1 {
		return false
	}
	n.stack = n.stack[:len(n.stack)-1]
	return true
}

// CanGoBack returns true if there's a screen to go back to.
func (n *Navigator) CanGoBack() bool {
	return len(n.stack) > 1
}

// Breadcrumb returns the navigation path as a string.
func (n *Navigator) Breadcrumb() string {
	parts := make([]string, len(n.stack))
	for i, e := range n.stack {
		parts[i] = e.Title
	}
	return strings.Join(parts, " > ")
}

// Navigation messages

// PushScreenMsg requests pushing a new screen.
type PushScreenMsg struct {
	Entry Entry
}

// PopScreenMsg requests going back one screen.
type PopScreenMsg struct{}

// PushScreen creates a command to push a screen.
func PushScreen(e Entry) tea.Cmd {
	return func() tea.Msg { return PushScreenMsg{Entry: e} }
}

// PopScreen creates a command to go back.
func PopScreen() tea.Cmd {
	return func() tea.Msg { return PopScreenMsg{} }
}
