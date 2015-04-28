package main

import (
	_ "Docker-admin/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
