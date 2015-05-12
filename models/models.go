package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/redis"
	logpkg "log"
	"os"
)

var rc redis.Client
var log *logpkg.Logger

func init() {
	rc.Addr = beego.AppConfig.String("redis_addr")
	log = logpkg.New(os.Stderr, "model", logpkg.Ltime|logpkg.Llongfile)
}
func RegistDB() {

	db_str := beego.AppConfig.String("sql_conn_str")
	orm.RegisterDataBase("default", "mysql", db_str)
	orm.SetMaxIdleConns("default", 10)

	// orm.RegisterModelWithPrefix("tb_", new(Blog))
	// orm.RegisterModelWithPrefix("tb_", new(Tag))
	// orm.RegisterModelWithPrefix("tb_", new(Review))
	orm.RegisterModel(new(Blog), new(Tag), new(Review), new(ConnLog))

	//start ORM debug
	orm.Debug = true
	//create table
	force := false
	verbose := true
	orm.RunSyncdb("default", force, verbose)
}
