package screens

import (
	"avro_cli/internal/app/registry"
	"avro_cli/internal/domain"
	"avro_cli/internal/tui/nav"
	"avro_cli/internal/tui/styles"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// SearchModel provides a fuzzy search command palette.
type SearchModel struct {
	query   string
	results []domain.CommandDescriptor
	cursor  int
	width   int
	height  int
}

// NewSearchModel creates the search screen.
func NewSearchModel() SearchModel {
	return SearchModel{
		results: registry.Global().All(),
	}
}

func (m SearchModel) Init() tea.Cmd { return nil }

func (m SearchModel) Update(msg tea.Msg) (SearchModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.results)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.results) > 0 {
				cmd := m.results[m.cursor]
				return m, nav.PushScreen(nav.Entry{
					Screen: nav.CommandDetailScreen,
					Title:  cmd.FullName(),
					Data:   cmd,
				})
			}
		case "backspace":
			if len(m.query) > 0 {
				m.query = m.query[:len(m.query)-1]
				m.updateResults()
			}
		default:
			if len(msg.String()) == 1 {
				m.query += msg.String()
				m.updateResults()
			}
		}
	}
	return m, nil
}

func (m *SearchModel) updateResults() {
	if m.query == "" {
		m.results = registry.Global().All()
	} else {
		m.results = registry.Global().Search(m.query)
	}
	m.cursor = 0
}

func (m SearchModel) View() string {
	var b strings.Builder

	b.WriteString(styles.Subtitle.Render("Search Commands") + "\n\n")
	b.WriteString(fmt.Sprintf("  > %s_\n\n", m.query))

	maxVisible := 15
	if m.height > 0 {
		maxVisible = m.height - 8
		if maxVisible < 5 {
			maxVisible = 5
		}
	}

	start := 0
	if m.cursor >= maxVisible {
		start = m.cursor - maxVisible + 1
	}

	end := start + maxVisible
	if end > len(m.results) {
		end = len(m.results)
	}

	for i := start; i < end; i++ {
		cmd := m.results[i]
		line := fmt.Sprintf("%-24s %s", cmd.FullName(), styles.Description.Render(cmd.Description))
		if i == m.cursor {
			b.WriteString(styles.SelectedItem.Render(line))
		} else {
			b.WriteString(styles.NormalItem.Render(line))
		}
		b.WriteString("\n")
	}

	if len(m.results) == 0 {
		b.WriteString(styles.Description.Render("  No commands found") + "\n")
	}

	b.WriteString("\n")
	b.WriteString(styles.HelpStyle.Render("type to search | up/down: navigate | enter: select | esc: back"))

	return b.String()
}
