package main

import (
	"fmt"
	"friends/internal/commands"
	"friends/pkg/command"
	"os"
)

// TODO: General use packages: storage, cli
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {

	command.Load([]command.Command{
		commands.NewCmdInit(),
	})

	out := command.Run(os.Args[1:])
	if out.Err != nil {
		return out.Err
	}

	fmt.Println(out.Result)

	return nil
}
