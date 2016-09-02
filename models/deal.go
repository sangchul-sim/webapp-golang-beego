package models

import (
	_mysql "github.com/sangchul-sim/webapp-golang-beego/models/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TbDealInfo struct {
	DealId      int64  `orm:"pk"` // doesn't work without a PK, doesn't make sense, it's a fk...
	DealName    string `orm:"size(300)"`
	Price       int
	SalePrice   int
	SaleStartDt string `orm:"size(12)"`
	SaleEndDt   string `orm:"size(12)"`
}

func init() {
	err, _ := _mysql.MysqlConn()
	if err != nil {
		beego.Debug(err)
	}

	// This is compulsory if you use orm.QuerySeter for advanced query.
	// Otherwise you don’t need to do this if you use raw SQL query and map struct only.
	//
	// Register the Model you defined. The best practice is to have a single models.go file and register in it’s init function.
	// RegisterModel can register multiple models at the same time
	orm.RegisterModel(new(TbDealInfo))
}

// Raw SQL to query
func (deal *TbDealInfo) Deal(Id int64) (error, *TbDealInfo) {
	o := orm.NewOrm()
	o.Using("default")

	deal.DealId = Id
	query := `SELECT * FROM tb_deal_info WHERE deal_id = ?`

	err := o.Raw(query, deal.DealId).QueryRow(&deal)
	if err != nil {
		return err, deal
	}
	return nil, deal
}

// ORM to query
func (research *TbDealInfo) DealList(Limit int, Offset int, SearchKeyword string) (error, int64, []TbDealInfo) {
	o := orm.NewOrm()
	o.Using("default")

	var DealList []TbDealInfo

	qs := o.QueryTable("TbDealInfo")
	if SearchKeyword != "" {
		qs = qs.Filter("DealName__istartswith", SearchKeyword)
	}

	TotalCount, listErr := qs.OrderBy("-DealId").Limit(Limit, Offset).All(&DealList)

	if listErr != nil {
		return listErr, 0, nil
	}
	return nil, TotalCount, DealList
}
