package clite

import (
	"context"
	"os"
)

type App struct {
	Name        string
	Description string
	Cmds        []Cmd
}

func (a *App) Run(ctx context.Context) error {
	cmd := Cmd{
		Name:        a.Name,
		Description: a.Description,
		SubCmds:     a.Cmds,
		fullpath:    a.Name,
	}

	if len(os.Args) == 1 {
		cmd.printHelp()
		return nil
	}

	return cmd.Run(ctx, os.Args[1:])
}
