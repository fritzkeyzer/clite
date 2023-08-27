package clite

import (
	"fmt"

	"github.com/fritzkeyzer/conf"
)

type Cmd struct {
	Name        string
	Alias       string
	Description string

	// Flags should be a pointer to a struct (use field tags to define flags)
	// or nil if no flags are needed.
	// eg:
	// 	var cmdFlags struct {
	// 		Db      string `flag:"--db"    env:"DB"    comment:"database connection string"`
	// 		Hello   string `flag:"--hello" env:"HELLO" comment:"greets the user"`
	// 		Verbose bool   `flag:"-v"                  comment:"turns on verbose mode"`
	// 	}
	//
	// 	var cmd = clite.Cmd{
	// 		Flags: &cmdFlags,
	// 	}
	Flags    any
	Func     func() error // function to execute
	SubCmds  []Cmd
	fullpath string // populated at runtime to print the full path to the command
}

func (c *Cmd) Run(args []string) error {
	if c.SubCmds == nil {
		if checkForHelp() {
			c.printHelp()
			return nil
		}

		if c.Func == nil {
			return fmt.Errorf("no subcommands or Exec.Func defined")
		}

		return c.exec()
	}

	if len(args) == 0 {
		c.printHelp()
		return nil
	}

	subCmdName := args[0]

	for _, subCmd := range c.SubCmds {
		if subCmd.Name == subCmdName || subCmd.Alias == subCmdName {
			subCmd.fullpath = c.fullpath + " " + subCmdName
			return subCmd.Run(args[1:])
		}
	}

	c.printHelp()

	if checkForHelp() {
		return nil
	}

	return fmt.Errorf("invalid command: %q", subCmdName)
}

func (c *Cmd) exec() error {
	if c.Flags != nil {
		if err := conf.LoadEnv(c.Flags); err != nil {
			return fmt.Errorf("loading env: %v", err)
		}

		if err := conf.LoadFlags(c.Flags); err != nil {
			return fmt.Errorf("loading flags: %v", err)
		}
	}

	if err := c.Func(); err != nil {
		return err
	}

	return nil
}
