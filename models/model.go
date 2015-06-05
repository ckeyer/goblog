package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/redis"
	logpkg "log"
	"os"
)

var (
	rc      redis.Client
	log     *logpkg.Logger
	db_str  string
	force   = false // force create tables
	verbose = true  // show sql
)

func init() {
	// rc.Addr = beego.AppConfig.String("redis_addr")
	// db_str = beego.AppConfig.String("sql_conn_str")
	db_str = "root:root@/db_blog?charset=utf8"
	rc.Addr = "localhost:6379"
	log = logpkg.New(os.Stderr, "model", logpkg.Ltime|logpkg.Lshortfile)
}

func RegistDB() {
	orm.Debug = true

	orm.RegisterDataBase("default", "mysql", db_str)
	orm.SetMaxIdleConns("default", 10)
	orm.RegisterModel(new(Blog), new(Tag), new(Review), new(ConnLog), new(BlogTagRelation))

	orm.RunSyncdb("default", force, verbose)
}
