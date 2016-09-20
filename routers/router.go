package routers

import (
	"github.com/astaxie/beego"
	"github.com/sangchul-sim/webapp-golang-beego/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})

	beego.Router("/visitor/health_check",
		&controllers.MainController{},
		"get:HealthCheck")

	// Indicate ViewController.Main method to handle GET requests.
	beego.Router("/beego",
		&controllers.MainController{},
		"get:Beego")

	beego.Router("/beego_test",
		&controllers.MainController{},
		"get:Beego")

	beego.Router("/api/deal/:id([0-9]+)",
		&controllers.APIController{},
		"get:Deal")

	beego.Router("/api/deal/list",
		&controllers.APIController{},
		"get:DealList")

	beego.Router("/chat/home",
		&controllers.ChatController{},
		"get:Home")

	beego.Router("/chat/ws",
		&controllers.ChatController{},
		"get:Ws")
}
