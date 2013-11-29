package main

import (
	//	"fmt"
	"github.com/astaxie/beego"
	"github.com/fhbzyc/linuxbrowser/controllers"
)

func main() {

	beego.Router(`/`, &controllers.IndexController{})
	beego.Router(`/operation`, &controllers.OperationController{})
	beego.Router(`/:(.+)`, &controllers.IndexController{})
	//	beego.AutoRouter(&controllers.OperationController{})

	beego.SetStaticPath("/static", "static")

	beego.Run()
}
