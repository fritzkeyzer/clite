package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fritzkeyzer/clite"
)

/*
example
Some example commands

Usage:
  cli-example example <command> [flags]

Flags:
  --help  -h  Prints help for the command

Commands:
  a    Demonstrates using flags
  b    Error handling example
  c    Another nested command
  d    Demonstrates how you can create trees of commands
*/

var nestTopLevelCmd = clite.Cmd{
	Name:        "example",
	Description: "Some example commands",
	SubCmds: []clite.Cmd{
		nestedACmd,
		nestedBCmd,
		nestedCCmd,
		nestedDCmd,
	},
}

var nestedAOpts struct {
	Name    string `flag:"--name" comment:"A flag"`
	Age     int    `flag:"--age" comment:"A flag"`
	Verbose bool   `flag:"--age" comment:"A flag"`
}

// nestedACmd shows how flags of different types can easily be used eg:
//
//	$ cli example a --name='john' --age=42
//	2024/11/01 21:15:33 john is 42 years old
var nestedACmd = clite.Cmd{
	Name:        "a",
	Description: "Demonstrates using flags",
	Flags:       &nestedAOpts, // this lets clite parse your flags so you dont need to
	Func: func(ctx context.Context) error {
		// you can just use them
		if nestedAOpts.Verbose {
			log.Println("verbose mode enabled")
		}

		log.Printf("%s is %d years old", nestedAOpts.Name, nestedAOpts.Age)

		return nil
	},
}

var nestedBCmd = clite.Cmd{
	Name:        "b",
	Description: "Error handling example",
	Func: func(ctx context.Context) error {
		// this error is returned with app.Run(ctx)
		// you can handle fatal errors at the top level and no need to duplicate for each command
		return fmt.Errorf("something went wrong")
	},
}

var nestedCCmd = clite.Cmd{
	Name:        "c",
	Alias:       "",
	Description: "Another nested command",
	Flags:       nil,
	Func:        nil,
	SubCmds: []clite.Cmd{
		nestedDCmd,
	},
}

var nestedDCmd = clite.Cmd{
	Name:        "d",
	Description: "Demonstrates how you can create trees of commands",
}
