package controllers

import "github.com/astaxie/beego"

type MainController struct {
	BaseController
}

// @Title Health check
// @Description find object by id
// @Param	Id		path 	int64	true		"the DealID you want to get"
// @Success 200 {object} models.TbDealInfo
// @Failure 400 :id is empty
// @router /beego [get]
func (c *MainController) Beego() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.Render()
}

// @Title Health check
// @Description find object by id
// @Param	Id		path 	int64	true		"the DealID you want to get"
// @Success 200 {object} models.TbDealInfo
// @Failure 400 :id is empty
// @router /visitor/health_check [get]
func (c *MainController) HealthCheck() {
	AppVersion := beego.AppConfig.String("appversion")

	data := make(map[string]interface{})
	data["version"] = AppVersion

	c.Data["json"] = &data
	c.ServeJSON()
}
