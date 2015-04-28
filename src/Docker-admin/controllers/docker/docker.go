package docker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	Prot struct {
		IP          string `json:"IP"`
		PrivatePort int    `json:"PrivatePort"`
		PublicPort  int    `json:"PublicPort"`
		Type        string `json:"Type"`
	}

	Container struct {
		Command string   `json:"Command"`
		Created int      `json:"Created"`
		Id      string   `json:"Id"`
		Image   string   `json:"Image"`
		Names   []string `json:"Names"`
		Ports   []Prot   `json:"Ports"`
		Status  string   `json:"Status"`
	}

	InspectConfig struct {
		AttachStderr bool                `json:"AttachStderr"`
		AttachStdin  bool                `json:"AttachStdin"`
		AttachStdout bool                `json:"AttachStdout"`
		Cmd          []string            `json:"Cmd"`
		CpuShares    float64             `json:"CpuShares"`
		Cpuset       string              `json:"Cpuset"`
		Domainname   string              `json:"Domainname"`
		Entrypoint   []string            `json:"Entrypoint"`
		Env          []string            `json:"Env"`
		ExposedPorts map[string]struct{} `json:"ExposedPorts"`
		Hostname     string              `json:"Hostname"`
		Image        string              `json:"Image"`
		//MacAddress      string              `json:"MacAddress"`
		Memory          float64             `json:"Memory"`
		MemorySwap      float64             `json:"MemorySwap"`
		NetworkDisabled bool                `json:"NetworkDisabled"`
		OnBuild         []string            `json:"OnBuild"`
		OpenStdin       bool                `json:"OpenStdin"`
		PortSpecs       string              `json:"PortSpecs"`
		StdinOnce       bool                `json:"StdinOnce"`
		Tty             bool                `json:"Tty"`
		User            string              `json:"User"`
		Volumes         map[string]struct{} `json:"Volumes"`
		WorkingDir      string              `json:"WorkingDir"`
	}

	HostRestart struct {
		MaximumRetryCount string `json:"MaximumRetryCount"`
		Name              string `json:"Name"`
	}

	InspectNetworkPorts struct {
		HostIp   string `json:"HostIp"`
		HostPort string `json:"HostPort"`
	}

	InspectHost struct {
		Binds []string `json:"Binds"`
		//CapAdd          []string                         `json:"CapAdd"`
		//CapDrop         []string                         `json:"CapDrop"`
		ContainerIDFile string `json:"ContainerIDFile"`
		//Devices         []string                         `json:"Devices"`
		Dns       []string `json:"Dns"`
		DnsSearch []string `json:"DnsSearch"`
		//ExtraHosts      string                           `json:"ExtraHosts"`
		//IpcMode         string                           `json:"IpcMode"`
		Links        []string                         `json:"Links"`
		LxcConf      []map[string]string              `json:"LxcConf"`
		NetworkMode  string                           `json:"NetworkMode"`
		PortBindings map[string][]InspectNetworkPorts `json:"PortBindings"`
		Privileged   bool                             `json:"Privileged"`
		//ReadonlyRootfs  bool                             `json:"ReadonlyRootfs"`
		PublishAllPorts bool `json:"PublishAllPorts"`
		//RestartPolicy   HostRestart                      `json:"RestartPolicy"`
		//SecurityOpt     string                           `json:"SecurityOpt"`
		VolumesFrom []string `json:"VolumesFrom"`
	}

	InspectNetwork struct {
		Bridge      string                           `json:"Bridge"`
		Gateway     string                           `json:"Gateway"`
		IPAddress   string                           `json:"IPAddress"`
		IPPrefixLen int                              `json:"IPPrefixLen"`
		MacAddress  string                           `json:"MacAddress"`
		PortMapping []string                         `json:"PortMapping"`
		Ports       map[string][]InspectNetworkPorts `json:"Ports"`
	}

	InspectState struct {
		//Error      string    `json:"Error"`
		ExitCode   int       `json:"ExitCode"`
		FinishedAt time.Time `json:"FinishedAt"`
		//OOMKilled  bool      `json:"OOMKilled"`
		Paused bool `json:"Paused"`
		Pid    int  `json:"Pid"`
		//Restarting bool      `json:"Restarting"`
		Running   bool      `json:"Running"`
		StartedAt time.Time `json:"StartedAt"`
		//Ghost      bool      `json:"Ghost"`
	}

	InspectContainer struct {
		//AppArmorProfile string            `json:"AppArmorProfile"`
		Args       []string       `json:"Args"`
		Config     *InspectConfig `json:"Config"`
		Created    string         `json:"Created"`
		Driver     string         `json:"Driver"`
		ExecDriver string         `json:"ExecDriver"`
		//ExecIDs         []string          `json:"ExecIDs"`
		HostConfig      *InspectHost   `json:"HostConfig"`
		HostnamePath    string         `json:"HostnamePath"`
		HostsPath       string         `json:"HostsPath"`
		Id              string         `json:"Id"`
		Image           string         `json:"Image"`
		MountLabel      string         `json:"MountLabel"`
		Name            string         `json:"Name"`
		NetworkSettings InspectNetwork `json:"NetworkSettings"`
		Path            string         `json:"Path"`
		ProcessLabel    string         `json:"ProcessLabel"`
		ResolvConfPath  string         `json:"ResolvConfPath"`
		//RestartCount    int               `json:"RestartCount"`
		State     InspectState      `json:"State"`
		Volumes   map[string]string `json:"Volumes"`
		VolumesRW map[string]bool   `json:"VolumesRW"`
	}
)

func GetHTML(url string) string {
	res, err := http.Get(url)
	if err != nil {
		return ""
	}
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	bodystr := string(body)
	return bodystr
}

func GetContainers(name string) ([]*Container, error) {
	var containers []*Container
	host, err := ReadHost(name)
	if err != nil {
		return nil, err
	}
	hosturl := fmt.Sprintf("http://%s:%d/containers/json?all=1", host.Ip, host.Port)
	body := GetHTML(hosturl)
	if err := json.Unmarshal([]byte(body), &containers); err != nil {
		return nil, err
	}
	return containers, nil
}

func GetInspectContain(ip string, port int, id string) (*InspectContainer, error) {
	var inspectContainer *InspectContainer
	url := fmt.Sprint("http://%s:%d/containers/%s/json", ip, port, id)
	body := GetHTML(url)
	if err := json.Unmarshal([]byte(body), &inspectContainer); err != nil {
		panic(err)
		return nil, err
	}
	return inspectContainer, nil
}
