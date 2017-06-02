package controllers

import "github.com/astaxie/beego"

type FrontpageController struct {
	beego.Controller
}

func (this *FrontpageController) Get() {
	this.TplName = "index.tpl"
	this.Render()
}
