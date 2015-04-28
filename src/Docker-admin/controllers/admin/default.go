package admin

import (
	"Docker-admin/controllers/docker"
	"github.com/astaxie/beego"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {
	this.Data["Website"] = "Docker Admin"
	this.Data["Email"] = "hidaen@gmail.com"
	this.Layout = "layout.html"
	this.TplNames = "index.html"
}

func (this *MainController) Container() {
	containers, _ := docker.ReadHosts()
	if hostname := this.Ctx.Input.Param(":hostname"); hostname != "" {
		if host, ok := containers[hostname]; ok {
			this.Data["Hostname"] = hostname
			containers, _ := docker.GetContainers(hostname)
			var up, exited []*docker.Container
			for k, v := range containers {
				if strings.HasPrefix(v.Status, "Up") {
					up = append(up, containers[k])
				} else {
					exited = append(exited, containers[k])
				}
			}
			this.Data["Up"] = up
			this.Data["Exited"] = exited
			if id := this.Ctx.Input.Param(":id"); id != "" {
				this.Data["Id"] = id
				inspectContainer, err := docker.GetInspectContain(host.Ip, host.Port, id)
				if err != nil {
					panic(inspectContainer)
				}
				this.Data["inspectContainer"] = inspectContainer
			} else {
				this.Data["Id"] = ""
			}
		} else {
			this.Data["Hostname"] = ""
		}
	} else {
		this.Data["Hostname"] = ""
	}
	this.Data["Hosts"] = containers
	this.Layout = "layout.html"
	this.TplNames = "container.html"
}
