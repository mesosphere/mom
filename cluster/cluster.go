package cluster

import (
	"github.com/mesosphere/mom/configuration"
	"github.com/mesosphere/mom/templates"
	"net/http"
  "bytes"
  "io/ioutil"
  "fmt"
)

const appPath string = "/v2/apps"

type Cluster struct {
	conf configuration.Configuration
}

func New(conf configuration.Configuration) *Cluster {
	return &Cluster{conf}
}

func (c *Cluster) Launch() error {
	// TODO(nnielsen): http post templated content to marathon endpoint.
	masterTemplate := templates.MasterTemplate{
		MasterAppId: "foobar",
	}
	masterJson := templates.FormatMaster(masterTemplate)

	println("Post: ", masterJson)

  url := c.conf.MarathonUrl + appPath

	var jsonStr = []byte(masterJson)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
    fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("response Body:", string(body))

  return nil
}
