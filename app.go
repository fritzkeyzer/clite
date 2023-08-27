package clite

import (
	"log"
	"os"
)

type App struct {
	Name        string
	Description string
	Cmds        []Cmd
}

func (a *App) Run() {
	cmd := Cmd{
		Name:        a.Name,
		Description: a.Description,
		SubCmds:     a.Cmds,
		fullpath:    a.Name,
	}

	if len(os.Args) == 1 {
		cmd.printHelp()
		return
	}

	err := cmd.Run(os.Args[1:])
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
}
