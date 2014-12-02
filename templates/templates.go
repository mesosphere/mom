package templates

import (
	"bytes"
	"log"
	"os"
	"path"
	"text/template"
)

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

const repoPath string = "src/github.com/mesosphere/mom"

func FormatMaster(masterTemplate MasterTemplate) string {
	goPath := os.Getenv("GOPATH")
	templatePath := path.Join(goPath, repoPath, "templates")

	masterPath := path.Join(templatePath, "mesos-master.json")
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
