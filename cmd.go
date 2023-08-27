package clite

import (
	"fmt"

	"github.com/fritzkeyzer/conf"
)

type Cmd struct {
	Name        string
	Alias       string
	Description string
	Params      any          // should be a pointer to a struct
	Func        func() error // function to execute
	SubCmds     []Cmd

	fullpath string // populated at runtime to print the full path to the command
}

func (c *Cmd) Run(args []string) error {
	if c.SubCmds == nil {
		if checkForHelp() {
			c.PrintHelp()
			return nil
		}

		if c.Func == nil {
			return fmt.Errorf("no subcommands or Exec.Func defined")
		}

		return c.exec()
	}

	if len(args) == 0 {
		c.PrintHelp()
		return nil
	}

	subCmdName := args[0]

	for _, subCmd := range c.SubCmds {
		if subCmd.Name == subCmdName || subCmd.Alias == subCmdName {
			subCmd.fullpath = c.fullpath + " " + subCmdName
			return subCmd.Run(args[1:])
		}
	}

	c.PrintHelp()

	if checkForHelp() {
		return nil
	}

	return fmt.Errorf("invalid command: %q", subCmdName)
}

func (c *Cmd) exec() error {
	if c.Params != nil {
		if err := conf.LoadEnv(c.Params); err != nil {
			return fmt.Errorf("loading env: %v", err)
		}

		if err := conf.LoadFlags(c.Params); err != nil {
			return fmt.Errorf("loading flags: %v", err)
		}
	}

	if err := c.Func(); err != nil {
		return err
	}

	return nil
}
