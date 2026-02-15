package cli

import (
	"avro_cli/internal/app/executor"
	"avro_cli/internal/app/registry"
	"avro_cli/internal/domain"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// BuildCobraTree creates cobra commands from the registry and attaches them to root.
func BuildCobraTree(root *cobra.Command, exec *executor.Executor) {
	reg := registry.Global()
	categoryCmds := make(map[string]*cobra.Command)

	for _, cat := range reg.Categories() {
		catCmd := &cobra.Command{
			Use:   cat.Name,
			Short: cat.Description,
		}
		categoryCmds[cat.Name] = catCmd
		root.AddCommand(catCmd)
	}

	for _, cmd := range reg.All() {
		leafCmd := buildLeafCommand(cmd, exec)
		if parent, ok := categoryCmds[cmd.Category.Name]; ok {
			parent.AddCommand(leafCmd)
		}
	}
}

func buildLeafCommand(desc domain.CommandDescriptor, exec *executor.Executor) *cobra.Command {
	cmd := &cobra.Command{
		Use:     buildUse(desc),
		Short:   desc.Description,
		Aliases: desc.Aliases,
		RunE: func(c *cobra.Command, args []string) error {
			flags := make(map[string]string)
			for _, f := range desc.Flags {
				val, _ := c.Flags().GetString(f.Name)
				if val != "" {
					flags[f.Name] = val
				}
			}

			result := exec.Run(desc, args, flags)
			if !result.IsOk() {
				fmt.Fprintln(os.Stderr, "Error:", result.Err())
				os.Exit(1)
			}
			if output := result.Value(); output != "" {
				fmt.Println(output)
			}
			return nil
		},
	}

	for _, f := range desc.Flags {
		cmd.Flags().StringP(f.Name, f.Short, f.Default, f.Description)
	}

	return cmd
}

func buildUse(desc domain.CommandDescriptor) string {
	use := desc.Name
	for _, a := range desc.Args {
		if a.Required {
			use += " <" + a.Name + ">"
		} else {
			use += " [" + a.Name + "]"
		}
	}
	return use
}
