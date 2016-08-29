package mysql

import (
	"database/sql"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var session *sql.DB
var connError error

func MysqlConn() (*sql.DB, error) {
	if connError != nil {
		return nil, connError
	}
	return session, nil
}

func init() {
	url := beego.AppConfig.String("mysql::url")

	beego.Debug("TypeOf(url)", reflect.TypeOf(url))
	beego.Debug("url", url)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	// ORM must register a database with alias default.
	err := orm.RegisterDataBase("default", "mysql", url)
	if err != nil {
		beego.Debug("error", err)

		connError = err
	}

	// Print out SQL query in debugging mode
	// http://beego.me/docs/mvc/model/orm.md
	RunMode := beego.AppConfig.String("runmode")
	if RunMode == "dev" {
		orm.Debug = true
	}
}
