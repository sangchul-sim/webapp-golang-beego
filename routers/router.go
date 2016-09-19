package routers

import (
	"github.com/astaxie/beego"
	"github.com/sangchul-sim/webapp-golang-beego/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSRouter("/visitor/health_check", &controllers.MainController{}, "get:HealthCheck"),

		//beego.NSNamespace("/visitor",
		//	beego.NSInclude(
		//		&controllers.MainController{},
		//	),
		//),
		beego.NSNamespace("/api/deal",
			beego.NSInclude(
				&controllers.APIController{},
			),
		),
	)
	beego.AddNamespace(ns)

	//
	//ns3 := beego.NewNamespace("/chat",
	//	beego.NSNamespace("/home",
	//		beego.NSInclude(
	//			&controllers.ChatController{},
	//		),
	//	),
	//	beego.NSNamespace("/ws",
	//		beego.NSInclude(
	//			&controllers.ChatController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns3)

	//beego.Router("/visitor/health_check",
	//	&controllers.MainController{},
	//	"get:HealthCheck")

	// Indicate ViewController.Main method to handle GET requests.
	//beego.Router("/beego",
	//	&controllers.MainController{},
	//	"get:Beego")
	//
	//beego.Router("/beego_test",
	//	&controllers.MainController{},
	//	"get:Beego")
	//
	//beego.Router("/api/deal/:id([0-9]+)",
	//	&controllers.APIController{},
	//	"get:Deal")
	//
	//beego.Router("/api/deal/list",
	//	&controllers.APIController{},
	//	"get:DealList")
	//
	//beego.Router("/chat/home",
	//	&controllers.ChatController{},
	//	"get:Home")
	//
	//beego.Router("/chat/ws",
	//	&controllers.ChatController{},
	//	"get:Ws")
}
