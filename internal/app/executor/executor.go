package executor

import (
	"avro_cli/internal/domain"
	"fmt"
)

// Executor validates arguments and runs a command action.
type Executor struct {
	Shell domain.ShellRunner
	FS    domain.FileSystem
	HTTP  domain.HTTPClient
}

// New creates an executor with the given infrastructure dependencies.
func New(shell domain.ShellRunner, fs domain.FileSystem, http domain.HTTPClient) *Executor {
	return &Executor{Shell: shell, FS: fs, HTTP: http}
}

// Run validates the provided args/flags against the command descriptor and executes the action.
func (e *Executor) Run(cmd domain.CommandDescriptor, args []string, flags map[string]string) domain.Result[string] {
	resolved, err := e.resolveArgs(cmd, args)
	if err != nil {
		return domain.Fail[string](err)
	}

	resolvedFlags := e.resolveFlags(cmd, flags)

	ctx := domain.CommandContext{
		Args:  resolved,
		Flags: resolvedFlags,
		Shell: e.Shell,
		FS:    e.FS,
		HTTP:  e.HTTP,
	}

	return cmd.Action(ctx)
}

func (e *Executor) resolveArgs(cmd domain.CommandDescriptor, provided []string) (map[string]string, error) {
	resolved := make(map[string]string)

	for i, def := range cmd.Args {
		if i < len(provided) {
			resolved[def.Name] = provided[i]
		} else if def.Required {
			return nil, &domain.ValidationError{
				Field:   def.Name,
				Message: fmt.Sprintf("required argument %q is missing", def.Name),
			}
		} else if def.Default != "" {
			resolved[def.Name] = def.Default
		}
	}

	return resolved, nil
}

func (e *Executor) resolveFlags(cmd domain.CommandDescriptor, provided map[string]string) map[string]string {
	resolved := make(map[string]string)

	for _, def := range cmd.Flags {
		if val, ok := provided[def.Name]; ok {
			resolved[def.Name] = val
		} else if def.Default != "" {
			resolved[def.Name] = def.Default
		}
	}

	return resolved
}
