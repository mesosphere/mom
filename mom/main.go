package main

import (
	"github.com/codegangsta/cli"
	"github.com/mesosphere/mom/configuration"
	"github.com/mesosphere/mom/cluster"
	"log"
	"os"
)

func main() {
	conf, err := configuration.Parse()
	if err != nil {
		log.Fatal("Could not parse configuration: ", err)
	}

	app := cli.NewApp()
	app.Name = "mom"
	app.Usage = "Mesos on Mesos cluster testing"
	app.Commands = []cli.Command{
		{
			Name:  "cluster",
			Usage: "Cluster utilies",
			Subcommands: []cli.Command{
				{
					Name:  "launch",
					Usage: "",
					Action: func(c *cli.Context) {
            cl := cluster.New(conf)
            cl.Launch()
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
