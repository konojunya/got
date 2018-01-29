package main

import (
	"os"

	"github.com/konojunya/got/command"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "got"
	app.Usage = "input your terminal `got new`!"

	app.Commands = []cli.Command{
		{
			Name:   "new",
			Usage:  "new go template!",
			Action: command.New,
		},
	}

	app.Run(os.Args)
}
