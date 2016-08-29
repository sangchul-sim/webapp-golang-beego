package main

import (
	"github.com/sangchul-sim/webapp-golang-beego/controllers"
	_ "github.com/sangchul-sim/webapp-golang-beego/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.EnableGzip = true
	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"}
	beego.BConfig.WebConfig.TemplateLeft = "{#"
	beego.BConfig.WebConfig.TemplateRight = "#}"
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.Log.AccessLogs = true
	beego.BConfig.Log.FileLineNum = true

	// access log
	beego.SetLogger("file", `{"filename":"logs/app-access.log", "rotate":true, "daily":true}`)

	// Enable XSRF
	beego.BConfig.WebConfig.EnableXSRF = true

	// XSRFKEY
	// XSRF key
	beego.BConfig.WebConfig.XSRFKey = "6QO07n7gTbZjAzN7m7ZBjL0ZCqL0o4EY"

	// XSRFExpire
	// XSRF expire time, 0 by default.
	beego.BConfig.WebConfig.XSRFExpire = 3600

	// beego.ErrorHandler("404", page_not_found)
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
