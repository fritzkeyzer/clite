package main

import (
	"context"
	"log"

	"github.com/fritzkeyzer/clite"
)

// $ cli-example some-cmd --db=123 --hello=world
// 2023/08/27 14:03:18 db: 123
// 2023/08/27 14:03:18 hello: world

var someCmdFlags struct {
	Db      string `flag:"--db"    env:"DB"    comment:"database connection string"`
	Hello   string `flag:"--hello" env:"HELLO" comment:"greets the user"`
	Verbose bool   `flag:"-v"                  comment:"turns on verbose mode"`
}

var someCmd = clite.Cmd{
	Name:        "some-cmd",
	Description: "Does something interesting",
	Flags:       &someCmdFlags,
	Func: func(ctx context.Context) error {
		if someCmdFlags.Verbose {
			log.Println("verbose output")
		}

		log.Println("db:", someCmdFlags.Db)
		log.Println("hello:", someCmdFlags.Hello)

		return nil
	},
}
