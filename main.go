package main

import (
	//	"fmt"
	"browser/controllers"
	"github.com/astaxie/beego"
)

func main() {

	beego.Router(`/`, &controllers.IndexController{})
	beego.Router(`/:(.+)`, &controllers.IndexController{})
	beego.Router(`/operation`, &controllers.OperationController{})
	//	beego.AutoRouter(&controllers.OperationController{})

	beego.SetStaticPath("/static", "static")

	beego.Run()
}
