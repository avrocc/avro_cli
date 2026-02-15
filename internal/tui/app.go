package tui

import (
	"avro_cli/internal/app/executor"
	"avro_cli/internal/app/registry"
	"avro_cli/internal/domain"
	"avro_cli/internal/infra/fs"
	"avro_cli/internal/infra/net"
	"avro_cli/internal/infra/shell"
	"avro_cli/internal/tui/components"
	"avro_cli/internal/tui/nav"
	"avro_cli/internal/tui/screens"
	"avro_cli/internal/tui/styles"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// appModel is the root Bubble Tea model that manages screen routing.
type appModel struct {
	nav    *nav.Navigator
	exec   *executor.Executor
	width  int
	height int

	// Screen models
	home     screens.HomeModel
	category screens.CategoryModel
	detail   screens.CommandDetailModel
	search   screens.SearchModel
}

func newAppModel() appModel {
	exec := executor.New(shell.New(), fs.New(), net.New())

	return appModel{
		nav:  nav.New(),
		exec: exec,
		home: screens.NewHomeModel(),
	}
}

func (m appModel) Init() tea.Cmd {
	return nil
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			current := m.nav.Current().Screen
			// Don't quit if typing in search or command detail
			if current == nav.SearchScreen || current == nav.CommandDetailScreen {
				break
			}
			if !m.nav.CanGoBack() {
				return m, tea.Quit
			}
		case "/":
			current := m.nav.Current().Screen
			if current != nav.SearchScreen && current != nav.CommandDetailScreen {
				m.search = screens.NewSearchModel()
				m.nav.Push(nav.Entry{Screen: nav.SearchScreen, Title: "Search"})
				return m, nil
			}
		case "esc":
			current := m.nav.Current().Screen
			// Let search screen handle its own esc (standalone quit vs pop)
			if current == nav.SearchScreen {
				break
			}
			if m.nav.Pop() {
				return m, nil
			}
			return m, tea.Quit
		}

	case nav.PushScreenMsg:
		m.nav.Push(msg.Entry)
		switch msg.Entry.Screen {
		case nav.CategoryScreen:
			m.category = screens.NewCategoryModel(msg.Entry.Data.(string))
		case nav.CommandDetailScreen:
			cmd := msg.Entry.Data.(domain.CommandDescriptor)
			m.detail = screens.NewCommandDetailModel(cmd, m.exec)
		case nav.SearchScreen:
			m.search = screens.NewSearchModel()
		}
		return m, nil

	case nav.PopScreenMsg:
		m.nav.Pop()
		return m, nil
	}

	// Delegate to current screen
	var cmd tea.Cmd
	switch m.nav.Current().Screen {
	case nav.HomeScreen:
		m.home, cmd = m.home.Update(msg)
	case nav.CategoryScreen:
		m.category, cmd = m.category.Update(msg)
	case nav.CommandDetailScreen:
		m.detail, cmd = m.detail.Update(msg)
	case nav.SearchScreen:
		m.search, cmd = m.search.Update(msg)
	}

	return m, cmd
}

func (m appModel) View() string {
	var content string

	switch m.nav.Current().Screen {
	case nav.HomeScreen:
		content = m.home.View()
	case nav.CategoryScreen:
		content = m.category.View()
	case nav.CommandDetailScreen:
		content = m.detail.View()
	case nav.SearchScreen:
		content = m.search.View()
	}

	breadcrumb := styles.Breadcrumb.Render(m.nav.Breadcrumb())
	statusBar := components.StatusBar(
		fmt.Sprintf("%d commands", len(registry.Global().All())),
		m.width,
	)

	return breadcrumb + "\n" + content + "\n\n" + statusBar
}

// Run starts the interactive TUI.
func Run() error {
	p := tea.NewProgram(newAppModel(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}

// RunPalette starts the TUI in palette mode (search screen only).
func RunPalette() error {
	p := tea.NewProgram(newPaletteModel(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}

func newPaletteModel() appModel {
	exec := executor.New(shell.New(), fs.New(), net.New())

	return appModel{
		nav:    nav.NewWithInitial(nav.SearchScreen, "Palette"),
		exec:   exec,
		search: screens.NewPaletteSearchModel(exec),
	}
}
