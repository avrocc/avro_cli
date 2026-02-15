package domain

// Category groups related commands (e.g., "git", "system", "http").
type Category struct {
	Name        string
	Description string
	Icon        string
}

// ArgDef defines a single argument or flag for a command.
type ArgDef struct {
	Name        string
	Short       string // single-char shorthand for flags
	Description string
	Required    bool
	Default     string
	Type        ArgType
}

// ArgType enumerates supported argument types.
type ArgType int

const (
	ArgString ArgType = iota
	ArgBool
	ArgInt
)

// CommandContext carries resolved arguments and dependencies to a command action.
type CommandContext struct {
	Args  map[string]string
	Flags map[string]string
	Shell ShellRunner
	FS    FileSystem
	HTTP  HTTPClient
}

// CommandAction is the function signature every command implements.
type CommandAction func(ctx CommandContext) Result[string]

// CommandDescriptor fully describes a command for registration, CLI routing, and TUI rendering.
type CommandDescriptor struct {
	Category    Category
	Name        string
	Aliases     []string
	Description string
	Args        []ArgDef // positional
	Flags       []ArgDef // --flags
	Action      CommandAction
}

// FullName returns "category name" (e.g., "git clone").
func (c CommandDescriptor) FullName() string {
	return c.Category.Name + " " + c.Name
}
