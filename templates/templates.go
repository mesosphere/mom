package templates

import (
	"bytes"
	"log"
	"os"
	"path"
	"text/template"
)

const repoPath string = "src/github.com/mesosphere/mom"

type MasterTemplate struct {
	MesosDockerImage string
	MasterAppId      string
	MasterCount      int
	MasterCpus       float64
	MasterMem        int
	ZookeeperUrl     string
	QuorumSize       int
	MasterFlags      string
}

func FormatMaster(masterTemplate MasterTemplate) string {
	goPath := os.Getenv("GOPATH")
	templatePath := path.Join(goPath, repoPath, "templates")

	masterPath := path.Join(templatePath, "default", "mesos-master.json")
	tmpl, err := template.ParseFiles(masterPath)
	if err != nil {
		log.Fatal("Could not parse master template:", err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, masterTemplate)
	if err != nil {
		log.Fatal("Could not specialize master template:", err)
	}

	return buf.String()
}

type SlaveTemplate struct {
	MesosDockerImage string
	SlaveAppId       string
	SlaveCount       int
	SlaveCpus        float64
	SlaveMem         int
	ZookeeperUrl     string
	SlaveFlags       string
}

func FormatSlave(slaveTemplate SlaveTemplate) string {
	goPath := os.Getenv("GOPATH")
	templatePath := path.Join(goPath, repoPath, "templates")

	slavePath := path.Join(templatePath, "default", "mesos-slave.json")
	tmpl, err := template.ParseFiles(slavePath)
	if err != nil {
		log.Fatal("Could not parse slave template:", err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, slaveTemplate)
	if err != nil {
		log.Fatal("Could not specialize slave template:", err)
	}

	return buf.String()
}
