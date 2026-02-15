package cli

import (
	"avro_cli/internal/app/executor"
	"avro_cli/internal/tui"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// NewRootCommand creates the root cobra command with dual-mode dispatch.
func NewRootCommand(exec *executor.Executor) *cobra.Command {
	root := &cobra.Command{
		Use:     "avro",
		Short:   "Avro - personal dev toolbox",
		Long:    "A personal dev toolbox CLI with interactive TUI mode.\nRun 'avro' without arguments to launch the interactive TUI.",
		Version: Version,
		RunE: func(cmd *cobra.Command, args []string) error {
			// If stdin is not a terminal, show help instead of TUI
			if !term.IsTerminal(int(os.Stdin.Fd())) {
				return cmd.Help()
			}
			return tui.Run()
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	paletteCmd := &cobra.Command{
		Use:     "palette",
		Aliases: []string{"p"},
		Short:   "Launch command palette",
		Long:    "Launch a fuzzy-search command palette to find and execute any registered command.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return tui.RunPalette()
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.AddCommand(paletteCmd)
	root.SetVersionTemplate("avro {{.Version}}\n")
	BuildCobraTree(root, exec)
	return root
}
