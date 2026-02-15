package fs

import "os"

// LocalFS implements domain.FileSystem using the local file system.
type LocalFS struct{}

// New creates a new local file system adapter.
func New() *LocalFS { return &LocalFS{} }

func (f *LocalFS) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (f *LocalFS) WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func (f *LocalFS) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (f *LocalFS) ListDir(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(entries))
	for i, e := range entries {
		names[i] = e.Name()
	}
	return names, nil
}
