package controllers

import (
	"strconv"

	"github.com/sangchul-sim/webapp-golang-beego/models"
)

// APIController 는 /api 에 대한 controller 입니다.
type APIController struct {
	BaseController
}

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

	// data := make(map[string]interface{})
	// data["deal"] = &result

	c.Data["json"] = &result
	c.ServeJSON()
}

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
