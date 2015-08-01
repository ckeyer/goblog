package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/goblog/lib/logging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/redis"
	"time"
)

var (
	rc      redis.Client
	log     = logging.GetLogger()
	db_str  string
	force   = false // force create tables
	verbose = true  // show sql
)

func init() {
	db_str = "root:root@/db_blog?charset=utf8"
	rc.Addr = "localhost:6379"
}

func RegistDB() {
	orm.Debug = true
	orm.DefaultTimeLoc = time.UTC

	orm.RegisterDataBase("default", "mysql", db_str)
	orm.SetMaxIdleConns("default", 10)
	orm.RegisterModel(new(Blog), new(Tag), new(Review),
		new(ConnLog), new(BlogTagRelation), new(User),
		new(Profile), new(Message))

	err := orm.RunSyncdb("default", force, verbose)
	if err != nil {
		panic(err)
	}
}
