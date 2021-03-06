package cluster

import (
	"fmt"
	"github.com/mesosphere/mom/configuration"
	"github.com/mesosphere/mom/marathon"
	"github.com/mesosphere/mom/templates"
	"github.com/nu7hatch/gouuid"
	"path"
)

const masterLabel string = "master"
const slaveLabel string = "slave"

type Cluster struct {
	conf configuration.Configuration
}

func New(conf configuration.Configuration) *Cluster {
	return &Cluster{conf}
}

func (c *Cluster) Status(session string) error {
	m := marathon.New(c.conf.MarathonUrl)

	appId := path.Join(c.conf.AppPrefix, session, masterLabel)

	apps, err := m.GetApp(appId)
	if err != nil {
		return err
	}

	fmt.Printf("masters:\n")
	for _, task := range apps.App.Tasks {
		if len(task.Ports) > 0 {
			fmt.Printf("\t%s:%d\n", task.Host, task.Ports[0])
		}
	}

	return nil
}

func (c *Cluster) Destroy(session string) error {
	m := marathon.New(c.conf.MarathonUrl)

	masterAppId := path.Join(c.conf.AppPrefix, session, masterLabel)
	slaveAppId := path.Join(c.conf.AppPrefix, session, slaveLabel)

  // Make sure apps exist before trying to destroy a stray application.
	_, err := m.GetApp(masterAppId)
	if err != nil {
		return err
	}

	_, err = m.GetApp(slaveAppId)
	if err != nil {
		return err
	}


  // Now, destroy masters and slaves.
	err = m.DestroyApp(masterAppId)
	if err != nil {
		return err
	}

	err = m.DestroyApp(slaveAppId)
	if err != nil {
		return err
	}

  fmt.Printf("Cluster session %s destroyed\n", session)

	return nil
}

func (c *Cluster) Launch(image string) error {
	m := marathon.New(c.conf.MarathonUrl)

	u, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("Could not generate UUID: ", err)
	}

	session := u.String()

	fmt.Println("Launching cluster id: ", session)

	if c.conf.DockerHub != "" {
		image = path.Join(c.conf.DockerHub, image)
	}

	masterAppId := path.Join(c.conf.AppPrefix, session, masterLabel)

	zookeeperUrl := c.conf.Zookeeper + session

	fmt.Printf("\nmasters:\n")
	fmt.Printf("\tinstances:\t%d\n", c.conf.MasterCount)
	fmt.Printf("\tcpu:\t\t%f\n", c.conf.MasterCpu)
	fmt.Printf("\tmem:\t\t%d\n", c.conf.MasterMem)

	masterTemplate := templates.MasterTemplate{
		MesosDockerImage: image,
		MasterAppId:      masterAppId,
		MasterCount:      c.conf.MasterCount,
		MasterCpus:       c.conf.MasterCpu,
		MasterMem:        c.conf.MasterMem,
		ZookeeperUrl:     zookeeperUrl,
		QuorumSize:       c.conf.QuorumSize,
		MasterFlags:      c.conf.MasterFlags,
	}

	// TODO(nnielsen): Allow flags to overwrite defaults.
	masterJson := templates.FormatMaster(masterTemplate)

	slaveAppId := path.Join(c.conf.AppPrefix, session, slaveLabel)

	fmt.Printf("\nslaves:\n")
	fmt.Printf("\tinstances:\t%d\n", c.conf.SlaveCount)
	fmt.Printf("\tcpu:\t\t%f\n", c.conf.SlaveCpu)
	fmt.Printf("\tmem:\t\t%d\n", c.conf.SlaveMem)

	slaveTemplate := templates.SlaveTemplate{
		MesosDockerImage: image,
		SlaveAppId:       slaveAppId,
		SlaveCount:       c.conf.SlaveCount,
		SlaveCpus:        c.conf.SlaveCpu,
		SlaveMem:         c.conf.SlaveMem,
		ZookeeperUrl:     zookeeperUrl,
		SlaveFlags:       c.conf.SlaveFlags,
	}

	slaveJson := templates.FormatSlave(slaveTemplate)

	err = m.CreateApp(masterJson)
	if err != nil {
		return err
	}

	err = m.CreateApp(slaveJson)
	if err != nil {
		return err
	}

	return nil
}
