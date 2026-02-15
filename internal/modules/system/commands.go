package system

import (
	"avro_cli/internal/cli"
	"avro_cli/internal/domain"
	"context"
	"encoding/json"
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

var updateCmd = domain.CommandDescriptor{
	Category:    category,
	Name:        "update",
	Description: "Check for CLI updates from GitHub releases",
	Action: func(ctx domain.CommandContext) domain.Result[string] {
		const url = "https://api.github.com/repos/avrocc/avro_cli/releases/latest"
		headers := map[string]string{
			"Accept": "application/vnd.github+json",
		}

		status, body, err := ctx.HTTP.Get(context.Background(), url, headers)
		if err != nil {
			return domain.Fail[string](fmt.Errorf("failed to check updates: %w", err))
		}
		if status != 200 {
			return domain.Fail[string](fmt.Errorf("GitHub API returned HTTP %d", status))
		}

		var release struct {
			TagName     string `json:"tag_name"`
			HTMLURL     string `json:"html_url"`
			PublishedAt string `json:"published_at"`
		}
		if err := json.Unmarshal([]byte(body), &release); err != nil {
			return domain.Fail[string](fmt.Errorf("failed to parse release info: %w", err))
		}

		latest := strings.TrimPrefix(release.TagName, "v")
		current := cli.Version

		var sb strings.Builder
		fmt.Fprintf(&sb, "Current version: %s\n", current)
		fmt.Fprintf(&sb, "Latest version:  %s\n", latest)
		fmt.Fprintf(&sb, "Published:       %s\n", release.PublishedAt)

		if current == latest {
			sb.WriteString("\n✓ You are up to date!")
		} else {
			fmt.Fprintf(&sb, "\n⬆ Update available!\n%s", release.HTMLURL)
		}

		return domain.Ok(sb.String())
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
