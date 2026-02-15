package shell

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
)

// Runner implements domain.ShellRunner using os/exec.
type Runner struct{}

// New creates a new shell runner.
func New() *Runner { return &Runner{} }

// Run executes a command and returns its combined output.
func (r *Runner) Run(ctx context.Context, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if stderr.Len() > 0 {
			return "", &ShellError{Command: name, Output: strings.TrimSpace(stderr.String()), Cause: err}
		}
		return "", err
	}
	return strings.TrimRight(stdout.String(), "\n"), nil
}

// RunDir executes a command in a specific directory.
func (r *Runner) RunDir(ctx context.Context, dir string, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if stderr.Len() > 0 {
			return "", &ShellError{Command: name, Output: strings.TrimSpace(stderr.String()), Cause: err}
		}
		return "", err
	}
	return strings.TrimRight(stdout.String(), "\n"), nil
}

// ShellError wraps a command execution failure with stderr output.
type ShellError struct {
	Command string
	Output  string
	Cause   error
}

func (e *ShellError) Error() string {
	return e.Output
}

func (e *ShellError) Unwrap() error { return e.Cause }
