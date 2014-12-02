package main

import (
	"encoding/json"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"os/user"
	"path"
	"text/template"
)

const repoPath string = "src/github.com/mesosphere/mom"

type Configuration struct {
	DockerHub   string `json:"docker_hub"`
	Zookeeper   string `json:"zookeeper"`
	MarathonUrl string `json:"marathon_url"`
	AppPrefix   string `json:"app_prefix"`
	QuorumSize  int    `json:"quorum_size"`
	MasterCount int    `json:"master_count"`
	SlaveCount  int    `json:"slave_count"`
	// TODO(nnielsen): Take default master and slave resources.
}

type MasterTemplate struct {
	MesosDockerImage string
	MasterAppId      string
	MasterCount      string
	MasterCpus       string
	MasterMem        string
	ZookeeperUrl     string
	QuorumSize       string
	MasterFlags      string
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(usr.HomeDir + "/.mom.json")
	if err != nil {
		log.Fatal("Could not open configuration file: ", err)
	}

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("Could not parse configuration file:", err)
	}

	goPath := os.Getenv("GOPATH")
	templatePath := path.Join(goPath, repoPath, "templates")

  masterTemplate := MasterTemplate{
    MasterAppId: "foobar",
  }
	tmpl, err := template.New("master").ParseFiles(path.Join(templatePath, "mesos-master.json"))
  tmpl.Execute(os.Stdout, masterTemplate)

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
