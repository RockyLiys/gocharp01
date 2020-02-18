package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "baidu.com"
	c.Data["Email"] = "rockylee.200800@gmail.com"
	c.TplName = "index.html"
}
