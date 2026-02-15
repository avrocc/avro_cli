package git

import (
	"avro_cli/internal/domain"
	"context"
	"fmt"
)

var cloneCmd = domain.CommandDescriptor{
	Category:    category,
	Name:        "clone",
	Description: "Clone a git repository",
	Args: []domain.ArgDef{
		{Name: "url", Description: "Repository URL", Required: true},
		{Name: "dir", Description: "Target directory (optional)", Required: false},
	},
	Action: func(ctx domain.CommandContext) domain.Result[string] {
		url := ctx.Args["url"]
		if url == "" {
			return domain.Fail[string](&domain.ValidationError{Field: "url", Message: "repository URL is required"})
		}

		args := []string{"clone", url}
		if dir := ctx.Args["dir"]; dir != "" {
			args = append(args, dir)
		}

		output, err := ctx.Shell.Run(context.Background(), "git", args...)
		if err != nil {
			return domain.Fail[string](err)
		}
		if output == "" {
			return domain.Ok(fmt.Sprintf("Cloned %s successfully", url))
		}
		return domain.Ok(output)
	},
}

var statusCmd = domain.CommandDescriptor{
	Category:    category,
	Name:        "status",
	Aliases:     []string{"st"},
	Description: "Show git status",
	Action: func(ctx domain.CommandContext) domain.Result[string] {
		output, err := ctx.Shell.Run(context.Background(), "git", "status", "--short")
		if err != nil {
			return domain.Fail[string](err)
		}
		if output == "" {
			return domain.Ok("Working tree clean")
		}
		return domain.Ok(output)
	},
}

var logCmd = domain.CommandDescriptor{
	Category:    category,
	Name:        "log",
	Description: "Show recent git log",
	Flags: []domain.ArgDef{
		{Name: "count", Short: "n", Description: "Number of commits", Default: "10"},
	},
	Action: func(ctx domain.CommandContext) domain.Result[string] {
		count := ctx.Flags["count"]
		if count == "" {
			count = "10"
		}
		output, err := ctx.Shell.Run(context.Background(), "git", "log",
			"--oneline", "--graph", "--decorate", fmt.Sprintf("-n%s", count))
		if err != nil {
			return domain.Fail[string](err)
		}
		return domain.Ok(output)
	},
}

var branchCmd = domain.CommandDescriptor{
	Category:    category,
	Name:        "branch",
	Aliases:     []string{"br"},
	Description: "List git branches",
	Flags: []domain.ArgDef{
		{Name: "all", Short: "a", Description: "Show all branches including remotes", Default: ""},
	},
	Action: func(ctx domain.CommandContext) domain.Result[string] {
		args := []string{"branch"}
		if ctx.Flags["all"] != "" {
			args = append(args, "-a")
		}
		output, err := ctx.Shell.Run(context.Background(), "git", args...)
		if err != nil {
			return domain.Fail[string](err)
		}
		return domain.Ok(output)
	},
}
