package main

import (
	"github.com/sangchul-sim/webapp-golang-beego/controllers"
	_ "github.com/sangchul-sim/webapp-golang-beego/routers"

	"github.com/astaxie/beego"
)

func main() {
	// http://beego.me/docs/mvc/controller/config.md
	//
	// Enable Gzip or not, false by default.
	// If Gzip is enabled, the output of template will be compressed by Gzip or zlib according to
	beego.BConfig.EnableGzip = true

	// Set a list of file extensions.
	// Any static file with the extension in the list will support gzip compression.
	// It supports .css and .js by default.
	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"}
	beego.BConfig.WebConfig.TemplateLeft = "{#"
	beego.BConfig.WebConfig.TemplateRight = "#}"

	// Use auto render or not, true by default.
	// Should set it to false for API application as there is no need to render templates.
	beego.BConfig.WebConfig.AutoRender = false

	// Output access logs or not. It won’t output access logs under prod mode by default.
	beego.BConfig.Log.AccessLogs = true

	// Whether to print line number or not. Default is true. This config is not supported in config file.
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
