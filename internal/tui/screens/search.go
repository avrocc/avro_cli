package screens

import (
	"avro_cli/internal/app/executor"
	"avro_cli/internal/app/registry"
	"avro_cli/internal/domain"
	"avro_cli/internal/tui/nav"
	"avro_cli/internal/tui/styles"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SearchModel provides a fuzzy search command palette.
type SearchModel struct {
	query      string
	results    []registry.FuzzyResult
	cursor     int
	width      int
	height     int
	standalone bool // true = palette mode (esc quits, inline exec)
	exec       *executor.Executor

	// inline execution state (standalone only)
	executed bool
	output   string
	hasError bool
}

// NewSearchModel creates the search screen (embedded in TUI).
func NewSearchModel() SearchModel {
	return SearchModel{
		results: registry.Global().FuzzySearch(""),
	}
}

// NewPaletteSearchModel creates the search screen in standalone palette mode.
func NewPaletteSearchModel(exec *executor.Executor) SearchModel {
	return SearchModel{
		results:    registry.Global().FuzzySearch(""),
		standalone: true,
		exec:       exec,
	}
}

func (m SearchModel) Init() tea.Cmd { return nil }

func (m SearchModel) Update(msg tea.Msg) (SearchModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		// After inline execution, any key quits
		if m.executed {
			return m, tea.Quit
		}

		switch msg.String() {
		case "esc":
			if m.standalone {
				return m, tea.Quit
			}
			return m, nav.PopScreen()
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
				cmd := m.results[m.cursor].Command
				// Standalone + no args = execute inline
				if m.standalone && len(cmd.Args) == 0 && len(cmd.Flags) == 0 {
					m.executeInline(cmd)
					return m, nil
				}
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

func (m *SearchModel) executeInline(cmd domain.CommandDescriptor) {
	result := m.exec.Run(cmd, nil, nil)
	m.executed = true
	if result.IsOk() {
		m.output = result.Value()
		m.hasError = false
	} else {
		m.output = result.Err().Error()
		m.hasError = true
	}
}

func (m *SearchModel) updateResults() {
	m.results = registry.Global().FuzzySearch(m.query)
	m.cursor = 0
}

// highlight renders a command name with matched characters highlighted.
var matchStyle = lipgloss.NewStyle().Foreground(styles.Secondary).Bold(true)

func highlight(text string, matchedIndexes []int) string {
	if len(matchedIndexes) == 0 {
		return text
	}
	matched := make(map[int]bool, len(matchedIndexes))
	for _, idx := range matchedIndexes {
		matched[idx] = true
	}
	var b strings.Builder
	for i, ch := range text {
		if matched[i] {
			b.WriteString(matchStyle.Render(string(ch)))
		} else {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func (m SearchModel) View() string {
	var b strings.Builder

	// Show inline execution result
	if m.executed {
		b.WriteString(styles.Subtitle.Render("Command Palette") + "\n\n")
		if m.hasError {
			b.WriteString(styles.ErrorText.Render("Error: ") + m.output)
		} else {
			outputView := m.output
			if outputView == "" {
				outputView = "(no output)"
			}
			b.WriteString(styles.OutputBox.Render(outputView))
		}
		b.WriteString("\n\n")
		b.WriteString(styles.HelpStyle.Render("press any key to exit"))
		return b.String()
	}

	title := "Search Commands"
	if m.standalone {
		title = "Command Palette"
	}
	b.WriteString(styles.Subtitle.Render(title) + "\n\n")
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
		r := m.results[i]
		name := highlight(r.Command.FullName(), r.MatchedIndexes)
		// Pad highlighted name to fixed width
		nameLen := len(r.Command.FullName())
		padding := ""
		if nameLen < 24 {
			padding = strings.Repeat(" ", 24-nameLen)
		}
		line := name + padding + " " + styles.Description.Render(r.Command.Description)
		if i == m.cursor {
			b.WriteString(styles.SelectedItem.Render("") + line)
		} else {
			b.WriteString(styles.NormalItem.Render("") + line)
		}
		b.WriteString("\n")
	}

	if len(m.results) == 0 {
		b.WriteString(styles.Description.Render("  No commands found") + "\n")
	}

	b.WriteString("\n")
	helpText := "type to search | up/down: navigate | enter: select | esc: back"
	if m.standalone {
		helpText = "type to search | up/down: navigate | enter: select | esc: quit"
	}
	b.WriteString(styles.HelpStyle.Render(helpText))

	return b.String()
}
