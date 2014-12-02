package configuration

import (
	"encoding/json"
	"os"
	"os/user"
)

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

func Parse() (Configuration, error) {
	configuration := Configuration{}
	usr, err := user.Current()
	if err != nil {
		return configuration, err
	}
	file, err := os.Open(usr.HomeDir + "/.mom.json")
	if err != nil {
		return configuration, err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return configuration, err
	}

	return configuration, nil
}
