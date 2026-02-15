package domain

import "context"

// ShellRunner executes OS commands.
type ShellRunner interface {
	Run(ctx context.Context, name string, args ...string) (string, error)
	RunDir(ctx context.Context, dir string, name string, args ...string) (string, error)
}

// FileSystem provides basic file operations.
type FileSystem interface {
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte) error
	Exists(path string) bool
	ListDir(path string) ([]string, error)
}

// HTTPClient performs HTTP requests.
type HTTPClient interface {
	Get(ctx context.Context, url string, headers map[string]string) (int, string, error)
	Post(ctx context.Context, url string, body string, headers map[string]string) (int, string, error)
}
