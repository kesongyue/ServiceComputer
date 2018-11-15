package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller //这相当于继承beego里的Controller类
}

func (controller *MainController) Get() { //重写Get方法
	username := controller.Ctx.Input.Param(":name")      //获取路由信息
	controller.Ctx.WriteString("Helloword! " + username) // 没用beego的模板，直接往网页写东西
}

func main() {
	beego.Router("/cloudgo/:name", &MainController{}) //设置路由，传入controller处理函数
	beego.Run(":9000")                                //在9000端口上运行
}
