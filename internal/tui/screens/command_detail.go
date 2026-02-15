package screens

import (
	"avro_cli/internal/app/executor"
	"avro_cli/internal/domain"
	"avro_cli/internal/tui/styles"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// CommandDetailModel shows a command form, executes it, and displays output.
type CommandDetailModel struct {
	cmd      domain.CommandDescriptor
	exec     *executor.Executor
	fields   []fieldEntry
	cursor   int
	output   string
	hasError bool
	executed bool
	width    int
	height   int
}

type fieldEntry struct {
	def   domain.ArgDef
	value string
	isArg bool // true for positional args, false for flags
}

// NewCommandDetailModel creates the command detail screen.
func NewCommandDetailModel(cmd domain.CommandDescriptor, exec *executor.Executor) CommandDetailModel {
	var fields []fieldEntry
	for _, a := range cmd.Args {
		fields = append(fields, fieldEntry{def: a, value: a.Default, isArg: true})
	}
	for _, f := range cmd.Flags {
		fields = append(fields, fieldEntry{def: f, value: f.Default, isArg: false})
	}

	return CommandDetailModel{
		cmd:    cmd,
		exec:   exec,
		fields: fields,
	}
}

func (m CommandDetailModel) Init() tea.Cmd { return nil }

func (m CommandDetailModel) Update(msg tea.Msg) (CommandDetailModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		if m.executed {
			// After execution, any key goes back to form
			switch msg.String() {
			case "r":
				m.executed = false
				m.output = ""
				m.hasError = false
			}
			return m, nil
		}

		switch msg.String() {
		case "up", "shift+tab":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "tab":
			if m.cursor < len(m.fields)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.fields) == 0 || m.cursor >= len(m.fields) {
				m.execute()
			} else if m.cursor == len(m.fields)-1 {
				// On last field, enter executes
				m.execute()
			} else {
				m.cursor++
			}
		case "ctrl+r":
			m.execute()
		case "backspace":
			if len(m.fields) > 0 && m.cursor < len(m.fields) {
				f := &m.fields[m.cursor]
				if len(f.value) > 0 {
					f.value = f.value[:len(f.value)-1]
				}
			}
		default:
			if len(msg.String()) == 1 && len(m.fields) > 0 && m.cursor < len(m.fields) {
				m.fields[m.cursor].value += msg.String()
			}
		}
	}
	return m, nil
}

func (m *CommandDetailModel) execute() {
	args := make([]string, 0)
	flags := make(map[string]string)

	for _, f := range m.fields {
		if f.isArg {
			args = append(args, f.value)
		} else if f.value != "" {
			flags[f.def.Name] = f.value
		}
	}

	result := m.exec.Run(m.cmd, args, flags)
	m.executed = true
	if result.IsOk() {
		m.output = result.Value()
		m.hasError = false
	} else {
		m.output = result.Err().Error()
		m.hasError = true
	}
}

func (m CommandDetailModel) View() string {
	var b strings.Builder

	b.WriteString(styles.Subtitle.Render(m.cmd.FullName()) + "\n")
	b.WriteString(styles.Description.Render(m.cmd.Description) + "\n\n")

	if len(m.fields) == 0 {
		b.WriteString(styles.Description.Render("No arguments required") + "\n")
	} else {
		for i, f := range m.fields {
			label := f.def.Name
			if f.def.Required {
				label += "*"
			}
			prefix := "  "
			if i == m.cursor && !m.executed {
				prefix = "> "
			}

			val := f.value
			if i == m.cursor && !m.executed {
				val += "_"
			}
			if val == "" && f.def.Default != "" {
				val = styles.Description.Render("(" + f.def.Default + ")")
			}

			line := fmt.Sprintf("%s%-16s %s", prefix, label+":", val)
			if i == m.cursor && !m.executed {
				b.WriteString(styles.SelectedItem.Render(line))
			} else {
				b.WriteString(styles.NormalItem.Render(line))
			}
			b.WriteString("\n")
		}
	}

	if m.executed {
		b.WriteString("\n")
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
		b.WriteString(styles.HelpStyle.Render("r: run again | esc: back"))
	} else {
		b.WriteString("\n")
		b.WriteString(styles.HelpStyle.Render("tab/shift+tab: navigate | ctrl+r: run | esc: back"))
	}

	return b.String()
}
