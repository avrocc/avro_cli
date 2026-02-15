package system

import (
	"avro_cli/internal/domain"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var infoCmd = domain.CommandDescriptor{
	Category:    category,
	Name:        "info",
	Description: "Show system information (OS, arch, Go version)",
	Action: func(ctx domain.CommandContext) domain.Result[string] {
		hostname, _ := os.Hostname()
		home, _ := os.UserHomeDir()
		info := fmt.Sprintf(
			"OS:       %s\nArch:     %s\nCPUs:     %d\nGo:       %s\nHostname: %s\nHome:     %s",
			runtime.GOOS, runtime.GOARCH, runtime.NumCPU(),
			runtime.Version(), hostname, home,
		)
		return domain.Ok(info)
	},
}

var envCmd = domain.CommandDescriptor{
	Category:    category,
	Name:        "env",
	Description: "List environment variables (optionally filtered by prefix)",
	Args: []domain.ArgDef{
		{Name: "filter", Description: "Filter prefix (optional)", Required: false},
	},
	Action: func(ctx domain.CommandContext) domain.Result[string] {
		filter := ctx.Args["filter"]
		var lines []string
		for _, e := range os.Environ() {
			if filter == "" || strings.HasPrefix(strings.ToUpper(e), strings.ToUpper(filter)) {
				lines = append(lines, e)
			}
		}
		if len(lines) == 0 {
			return domain.Ok("No matching environment variables found")
		}
		return domain.Ok(strings.Join(lines, "\n"))
	},
}

var pathCmd = domain.CommandDescriptor{
	Category:    category,
	Name:        "path",
	Description: "List PATH entries, one per line",
	Action: func(ctx domain.CommandContext) domain.Result[string] {
		pathVar := os.Getenv("PATH")
		separator := ":"
		if runtime.GOOS == "windows" {
			separator = ";"
		}
		entries := strings.Split(pathVar, separator)
		return domain.Ok(strings.Join(entries, "\n"))
	},
}
