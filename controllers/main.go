package controllers

import "github.com/astaxie/beego"

type MainController struct {
	BaseController
}

func (c *MainController) Beego() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.Render()
}

func (c *MainController) HealthCheck() {
	AppVersion := beego.AppConfig.String("appversion")

	data := make(map[string]interface{})
	data["version"] = AppVersion

	c.Data["json"] = &data
	c.ServeJSON()
}
