package controllers

import (
	"github.com/astaxie/beego"
)

// MainController => Start Page Controller
type MainController struct {
	beego.Controller
}

// Get => get start page
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

// AddController =>
type AddController struct {
	beego.Controller
}

// Get =>
func (c *AddController) Get() {
	c.Data["content"] = "value"
	c.Layout = "admin/layout.html"
	c.TplName = "admin/add.tpl"
}
