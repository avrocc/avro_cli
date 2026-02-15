package domain

import "fmt"

// CommandNotFoundError indicates a command was not found in the registry.
type CommandNotFoundError struct {
	Name string
}

func (e *CommandNotFoundError) Error() string {
	return fmt.Sprintf("command not found: %s", e.Name)
}

// ValidationError indicates invalid arguments or flags.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error for %q: %s", e.Field, e.Message)
}

// ExecutionError wraps errors from command execution.
type ExecutionError struct {
	Command string
	Cause   error
}

func (e *ExecutionError) Error() string {
	return fmt.Sprintf("execution error in %q: %v", e.Command, e.Cause)
}

func (e *ExecutionError) Unwrap() error { return e.Cause }
