package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

var (
	tplName = "error.html"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error400(args ...string) {
	c.Data["StatusCode"] = http.StatusBadRequest
	c.Data["StatusText"] = http.StatusText(http.StatusBadRequest)
	c.TplName = tplName
	c.Render()
}

func (c *ErrorController) Error401(args ...string) {
	c.Data["StatusCode"] = http.StatusUnauthorized
	c.Data["StatusText"] = http.StatusText(http.StatusUnauthorized)
	c.TplName = tplName
	c.Render()
}

func (c *ErrorController) Error402(args ...string) {
	c.Data["StatusCode"] = http.StatusPaymentRequired
	c.Data["StatusText"] = http.StatusText(http.StatusPaymentRequired)
	c.TplName = tplName
	c.Render()
}

func (c *ErrorController) Error403(args ...string) {
	c.Data["StatusCode"] = http.StatusForbidden
	c.Data["StatusText"] = http.StatusText(http.StatusForbidden)
	c.TplName = tplName
	c.Render()
}

func (c *ErrorController) Error404(args ...string) {
	c.Data["StatusCode"] = http.StatusNotFound
	c.Data["StatusText"] = http.StatusText(http.StatusNotFound)
	c.TplName = tplName
	c.Render()
}

func (c *ErrorController) Error405() {
	c.Data["StatusCode"] = http.StatusMethodNotAllowed
	c.Data["StatusText"] = http.StatusText(http.StatusMethodNotAllowed)
	c.TplName = tplName
	c.Render()
}

func (c *ErrorController) Error500() {
	c.Data["StatusCode"] = http.StatusInternalServerError
	c.Data["StatusText"] = http.StatusText(http.StatusInternalServerError)
	c.TplName = tplName
	c.Render()
}

func (c *ErrorController) ErrorDb() {
	c.Data["StatusCode"] = http.StatusInternalServerError
	c.Data["StatusText"] = "database is now down"
	c.TplName = tplName
	c.Render()
}
