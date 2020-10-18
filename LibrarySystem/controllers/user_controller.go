package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Data["username"] = "爱白菜的小昆虫"
	this.Data["age"] = "1111"
	this.Data["pwd"] = "asdfg"
	this.TplName = "user.html"
}
