package components

import (
	"avro_cli/internal/tui/styles"

	tea "github.com/charmbracelet/bubbletea"
)

// ConfirmResult is sent when the user answers the confirm dialog.
type ConfirmResult struct {
	Confirmed bool
}

// ConfirmModel is a yes/no confirmation dialog.
type ConfirmModel struct {
	Message string
	yes     bool
}

// NewConfirmModel creates a confirm dialog.
func NewConfirmModel(message string) ConfirmModel {
	return ConfirmModel{Message: message, yes: true}
}

func (m ConfirmModel) Init() tea.Cmd { return nil }

func (m ConfirmModel) Update(msg tea.Msg) (ConfirmModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h", "tab":
			m.yes = !m.yes
		case "right", "l":
			m.yes = !m.yes
		case "y":
			return m, func() tea.Msg { return ConfirmResult{Confirmed: true} }
		case "n":
			return m, func() tea.Msg { return ConfirmResult{Confirmed: false} }
		case "enter":
			return m, func() tea.Msg { return ConfirmResult{Confirmed: m.yes} }
		}
	}
	return m, nil
}

func (m ConfirmModel) View() string {
	yesLabel := "  Yes  "
	noLabel := "  No  "

	if m.yes {
		yesLabel = styles.SelectedItem.Render("[Yes]")
		noLabel = styles.NormalItem.Render(" No ")
	} else {
		yesLabel = styles.NormalItem.Render(" Yes ")
		noLabel = styles.SelectedItem.Render("[No]")
	}

	return m.Message + "\n\n" + yesLabel + "    " + noLabel
}
