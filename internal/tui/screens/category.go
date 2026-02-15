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

// CategoryModel displays commands within a category.
type CategoryModel struct {
	category string
	commands []domain.CommandDescriptor
	cursor   int
	width    int
	height   int
}

// NewCategoryModel creates a category screen for the given category name.
func NewCategoryModel(category string) CategoryModel {
	return CategoryModel{
		category: category,
		commands: registry.Global().ByCategory(category),
	}
}

func (m CategoryModel) Init() tea.Cmd { return nil }

func (m CategoryModel) Update(msg tea.Msg) (CategoryModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.commands)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.commands) > 0 {
				cmd := m.commands[m.cursor]
				return m, nav.PushScreen(nav.Entry{
					Screen: nav.CommandDetailScreen,
					Title:  cmd.Name,
					Data:   cmd,
				})
			}
		}
	}
	return m, nil
}

func (m CategoryModel) View() string {
	var b strings.Builder

	b.WriteString(styles.Subtitle.Render(m.category) + "\n\n")

	for i, cmd := range m.commands {
		line := fmt.Sprintf("%-16s %s", cmd.Name, styles.Description.Render(cmd.Description))
		if i == m.cursor {
			b.WriteString(styles.SelectedItem.Render(line))
		} else {
			b.WriteString(styles.NormalItem.Render(line))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(styles.HelpStyle.Render("j/k: navigate | enter: select | esc: back | /: search | q: quit"))

	return b.String()
}
