package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	Primary   = lipgloss.Color("#7C3AED")
	Secondary = lipgloss.Color("#06B6D4")
	Success   = lipgloss.Color("#10B981")
	Warning   = lipgloss.Color("#F59E0B")
	Error     = lipgloss.Color("#EF4444")
	Muted     = lipgloss.Color("#6B7280")
	BgDark    = lipgloss.Color("#1F2937")

	// Text styles
	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(Primary).
		MarginBottom(1)

	Subtitle = lipgloss.NewStyle().
			Foreground(Secondary).
			Bold(true)

	Description = lipgloss.NewStyle().
			Foreground(Muted)

	ErrorText = lipgloss.NewStyle().
			Foreground(Error).
			Bold(true)

	SuccessText = lipgloss.NewStyle().
			Foreground(Success)

	// Layout
	Breadcrumb = lipgloss.NewStyle().
			Foreground(Muted).
			MarginBottom(1)

	StatusBar = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(BgDark).
			Padding(0, 1)

	// List items
	SelectedItem = lipgloss.NewStyle().
			Foreground(Primary).
			Bold(true).
			PaddingLeft(2)

	NormalItem = lipgloss.NewStyle().
			PaddingLeft(2)

	// Containers
	OutputBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Muted).
			Padding(1, 2).
			MarginTop(1)

	HelpStyle = lipgloss.NewStyle().
			Foreground(Muted).
			MarginTop(1)
)
