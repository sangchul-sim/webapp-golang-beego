package controllers

import (
	"strconv"

	"github.com/sangchul-sim/webapp-golang-beego/models"
)

// APIController 는 /api 에 대한 controller 입니다.
type APIController struct {
	BaseController
}

// @Title Deal
// @Description find object by id
// @Param	Id		path 	string	true		"the DealID you want to get"
// @Success 200 {object} models.TbDealInfo
// @Failure 400 :id is empty
// @router /:id [get]
func (c *APIController) Deal() {
	DealID := c.Ctx.Input.Param(":id")
	IntDealID, err := strconv.ParseInt(DealID, 10, 64)

	if err != nil {
		c.RetError(errInputData)
		return
	}

	DealModel := models.TbDealInfo{}
	errResult, result := DealModel.Deal(IntDealID)

	if errResult != nil {
		c.RetError(errNoUser)
		return
	}

	c.Data["json"] = &result
	c.ServeJSON()
}

// @Title DealList
// @Description get all deals
// @Success 200 {object} models.TbDealInfo
// @Failure 403 :DealId is empty
// @router / [get]
func (c *APIController) DealList() {
	SearchKeyword := c.Input().Get("search_keyword")
	Page := c.Input().Get("page")

	intPage, err := strconv.Atoi(Page)
	if err != nil {
		intPage = 1
	}

	Offset := (intPage - 1) * PagePer

	DealModel := models.TbDealInfo{}

	errResult, _, result := DealModel.DealList(PagePer, Offset, SearchKeyword)
	if errResult != nil {
		c.RetError(errNoDealList)
		return
	}

	c.Data["json"] = &result
	c.ServeJSON()
}
