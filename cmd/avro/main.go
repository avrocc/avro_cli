package main

import (
	"avro_cli/internal/app/executor"
	"avro_cli/internal/cli"
	"avro_cli/internal/infra/fs"
	"avro_cli/internal/infra/net"
	"avro_cli/internal/infra/shell"
	"fmt"
	"os"

	// Auto-register all modules
	_ "avro_cli/internal/modules"
)

func main() {
	exec := executor.New(shell.New(), fs.New(), net.New())
	root := cli.NewRootCommand(exec)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
