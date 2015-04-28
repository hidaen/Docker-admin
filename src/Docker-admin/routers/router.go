package routers

import (
	"Docker-admin/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &admin.MainController{}, "*:Index")
	beego.Router("/container", &admin.MainController{}, "*:Container")
	beego.Router("/container/:hostname:string", &admin.MainController{}, "*:Container")
	beego.Router("/container/:hostname:string/:id:string", &admin.MainController{}, "*:Container")
}
