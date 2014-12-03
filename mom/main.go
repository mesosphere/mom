package main

import (
	"github.com/codegangsta/cli"
	"github.com/mesosphere/mom/cluster"
	"github.com/mesosphere/mom/configuration"
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
						dockerImage := ""
						if len(c.Args()) != 1 {
              log.Fatal("Launch requires docker image")
						}
						dockerImage = c.Args()[0]

						cl := cluster.New(conf)
						err = cl.Launch(dockerImage); if err != nil {
              log.Fatal(err)
            }
					},
				},
				{
					Name:  "status",
					Usage: "",
					Action: func(c *cli.Context) {
						dockerImage := ""
						if len(c.Args()) != 1 {
              log.Fatal("Status requires session id")
						}
						dockerImage = c.Args()[0]

						cl := cluster.New(conf)
						err = cl.Status(dockerImage); if err != nil {
              log.Fatal(err)
            }
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
