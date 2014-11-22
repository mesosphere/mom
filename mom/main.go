package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "mom"
	app.Usage = "Mesos on Mesos cluster testing"
	app.Commands = []cli.Command{
		{
			Name:  "image",
			Usage: "Mesos image utilities",
			Subcommands: []cli.Command{
				{
					Name:  "build",
					Usage: "",
					Action: func(c *cli.Context) {
						println("Not yet implemented")
					},
				},
				{
					Name:  "push",
					Usage: "",
					Action: func(c *cli.Context) {
						println("Not yet implemented")
					},
				},
			},
		},
		{
			Name:  "cluster",
			Usage: "Cluster utilies",
			Subcommands: []cli.Command{
				{
					Name:  "launch",
					Usage: "",
					Action: func(c *cli.Context) {
						println("Not yet implemented")
					},
				},
				{
					Name:  "scale",
					Usage: "",
					Action: func(c *cli.Context) {
						println("Not yet implemented")
					},
				},
				{
					Name:  "destroy",
					Usage: "",
					Action: func(c *cli.Context) {
						println("Not yet implemented")
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
