package main

import (
	_ "LibrarySystem/routers"
	"github.com/astaxie/beego"
)

func main() {
	// beego.SetViewsPath("front")// 将项目前端路径设为front目录下
	beego.Run()
}

