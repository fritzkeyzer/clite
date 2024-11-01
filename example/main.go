package main

import (
	"context"
	"log"

	"github.com/fritzkeyzer/clite"
)

/*
A simple example of how to use the clite package.

Usage:
  cli-example <command> [flags]

Flags:
  --help  -h  Prints help for the command

Commands:
  some-cmd    Does something interesting
  example     Some example commands
*/

func main() {
	app := clite.App{
		Name:        "cli-example",
		Description: "A simple example of how to use the clite package.",

		// both these commands are defined below
		Cmds: []clite.Cmd{
			someCmd,
			nestTopLevelCmd,
		},
	}

	// pass a custom context and handle errors in your main function
	err := app.Run(context.Background())
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
