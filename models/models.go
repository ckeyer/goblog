package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegistDB() {

	db_str := beego.AppConfig.String("sql_conn_str")
	orm.RegisterDataBase("default", "mysql", db_str)
	orm.SetMaxIdleConns("default", 10)

	orm.RegisterModelWithPrefix("tb_", new(Blog))
	orm.RegisterModelWithPrefix("tb_", new(Tag))
	orm.RegisterModelWithPrefix("tb_", new(Review))

	//start ORM debug
	orm.Debug = true
	//create table
	orm.RunSyncdb("default", true, true)

}
