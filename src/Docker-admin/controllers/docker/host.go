package docker

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var hostFile = "hosts.json"

type (
	Host struct {
		Ip   string `json:"Ip"`
		Port int    `json:"Port"`
	}
)

func ReadHosts() (map[string]Host, error) {
	hosts := map[string]Host{}
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "controllers/docker", hostFile)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &hosts); err != nil {
		return nil, err
	}
	return hosts, nil
}

func ReadHost(name string) (*Host, error) {
	hosts, err := ReadHosts()
	if err != nil {
		return nil, err
	}
	if h, ok := hosts[name]; ok {
		return &h, nil
	} else {
		return nil, errors.New("No Found " + name)
	}
}
