package routers

import (
	"github.com/astaxie/beego"
	"github.com/sangchul-sim/webapp-golang-beego/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})

	// Indicate ViewController.Main method to handle GET requests.
	beego.Router("/beego",
		&controllers.MainController{},
		"get:Get")

	beego.Router("/api/deal/:id([0-9]+)",
		&controllers.APIController{},
		"get:Deal")

	beego.Router("/api/deal/list",
		&controllers.APIController{},
		"get:DealList")
}
