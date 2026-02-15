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

// HomeModel displays the list of command categories.
type HomeModel struct {
	categories []domain.Category
	cursor     int
	width      int
	height     int
}

// NewHomeModel creates the home screen.
func NewHomeModel() HomeModel {
	return HomeModel{
		categories: registry.Global().Categories(),
	}
}

func (m HomeModel) Init() tea.Cmd { return nil }

func (m HomeModel) Update(msg tea.Msg) (HomeModel, tea.Cmd) {
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
			if m.cursor < len(m.categories)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.categories) > 0 {
				cat := m.categories[m.cursor]
				return m, nav.PushScreen(nav.Entry{
					Screen: nav.CategoryScreen,
					Title:  cat.Name,
					Data:   cat.Name,
				})
			}
		}
	}
	return m, nil
}

func (m HomeModel) View() string {
	var b strings.Builder

	b.WriteString(styles.Title.Render("avro") + "\n")
	b.WriteString(styles.Description.Render("Select a category") + "\n\n")

	for i, cat := range m.categories {
		icon := cat.Icon
		if icon == "" {
			icon = ">"
		}

		line := fmt.Sprintf("%s %s  %s", icon, cat.Name, styles.Description.Render(cat.Description))
		if i == m.cursor {
			b.WriteString(styles.SelectedItem.Render(line))
		} else {
			b.WriteString(styles.NormalItem.Render(line))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(styles.HelpStyle.Render("j/k: navigate | enter: select | /: search | q: quit"))

	return b.String()
}
