package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (controller *MainController) Get() {
	username := controller.Ctx.Input.Param(":name")
	controller.Ctx.WriteString("Helloword! " + username)
}

func main() {
	beego.Router("/cloudgo/:name", &MainController{})
	beego.Run(":9000")
}
