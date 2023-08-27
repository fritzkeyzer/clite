package clite

import (
	"bytes"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/olekukonko/tablewriter"
)

func checkForHelp() bool {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			return true
		}
	}

	return false
}

func (c *Cmd) printHelp() {
	fmt.Println(c.Name)
	if c.Description != "" {
		fmt.Println(c.Description)
	}
	fmt.Println()

	fmt.Println("Usage:")
	fmt.Println(" ", c.fullpath, "<command> [flags]")
	fmt.Println()

	fmt.Println("Flags:")
	// fmt.Println("  --help  -h  Prints help for the command")
	fmt.Println(getParamsHelp(c.Flags))

	if len(c.SubCmds) > 0 {

		fmt.Println("Commands:")

		tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		for _, subcmd := range c.SubCmds {
			tw.Write([]byte(fmt.Sprintf("\t%s\t%s\t%s\n", subcmd.Name, subcmd.Alias, subcmd.Description)))
		}
		tw.Flush()
	}
	fmt.Println()
}

func getParamsHelp(ptr any) string {
	buf := &bytes.Buffer{}
	tw := tablewriter.NewWriter(buf)
	tw.SetBorder(false)
	tw.SetColumnSeparator("")

	tw.Append([]string{"--help", "-h", "Prints help for the command"})

	if ptr == nil {
		tw.Render()
		return buf.String()
	}

	fields, err := paramFields(ptr)
	if err != nil {
		return "ERROR: paramFields(ptr): " + err.Error()
	}

	for _, field := range fields {
		flagName, flag := field.flagName()

		envVar, env := field.envVar()
		if env {
			envVar = "$" + envVar
		}

		comment := field.field.Tag.Get("comment")

		if !flag && !env {
			continue
		}

		tw.Append([]string{flagName, envVar, comment})
	}

	tw.Render()
	return buf.String()
}
