package main

import (
	"fmt"
	"log"

	"github.com/fritzkeyzer/clite"
)

/*
	cli-example
	A simple example of how to use the clite package.

	Usage:
	  cli-example <command> [flags]

	Params:
	  --help  -h  Prints help for the command

	Commands:
	  some-cmd      Does something interesting
	  demo-error    Demonstrates how to return an error
*/

func main() {
	app := clite.App{
		Name:        "cli-example",
		Description: "A simple example of how to use the clite package.",
		Cmds: []clite.Cmd{
			someCmd,
			demoErrorCmd,
		},
	}

	app.Run()

}

/*
	some-cmd
	Does something interesting

	Usage:
	  cli-example some-cmd <command> [flags]

	Params:
	  --help   -h      Prints help for the command
	  --db     $DB     database connection string
	  --hello  $HELLO  greets the user
	  -v               turns on verbose mode
*/

var SomeCfg struct {
	Db      string `flag:"--db"    env:"DB"    comment:"database connection string"`
	Hello   string `flag:"--hello" env:"HELLO" comment:"greets the user"`
	Verbose bool   `flag:"-v"                  comment:"turns on verbose mode"`
}

var someCmd = clite.Cmd{
	Name:        "some-cmd",
	Description: "Does something interesting",
	Params:      &SomeCfg,
	Func: func() error {
		if SomeCfg.Verbose {
			log.Println("verbose output")
		}

		log.Println("db:", SomeCfg.Db)
		log.Println("hello:", SomeCfg.Hello)

		return nil
	},
}

/*
	demo-error
	2023/08/22 21:06:13 ERROR: something went wrong
*/

var demoErrorCmd = clite.Cmd{
	Name:        "demo-error",
	Description: "Demonstrates how to return an error",
	Func: func() error {
		return fmt.Errorf("something went wrong")
	},
}