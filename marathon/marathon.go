package marathon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const appPath string = "v2/apps"

type Marathon struct {
	url    string
	client *http.Client
}

type Task struct {
	AppId string `json:"appId"`
	Id    string `json:"id"`
	Host  string `json:"host"`
	Ports []int  `json:"ports"`
}

type AppInfo struct {
	Tasks []Task `json:"tasks"`
}

type Apps struct {
	App AppInfo `json:"app"`
}

func New(url string) *Marathon {
	return &Marathon{
		url:    url,
		client: &http.Client{},
	}
}

func (m *Marathon) CreateApp(app string) error {
	url := strings.Join([]string{m.url, appPath}, "/")

	var jsonStr = []byte(app)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 201 {
		return fmt.Errorf("Could not launch cluster: %s %s", resp.Status, body)
	}

	// TODO(nnielsen): Log returned JSON
	_ = body

	return nil
}

func (m *Marathon) GetApp(appId string) (*Apps, error) {
	res := &Apps{}
	url := strings.Join([]string{m.url, appPath, appId}, "/")
	req, err := http.NewRequest("GET", url, nil)

	resp, err := m.client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return res, fmt.Errorf("Could not get app details: %s", resp.Status)
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (m *Marathon) DestroyApp(appId string) error {
	url := strings.Join([]string{m.url, appPath, appId}, "/")
	req, err := http.NewRequest("DELETE", url, nil)

	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Could not get app details: %s", resp.Status)
	}

	return nil
}
