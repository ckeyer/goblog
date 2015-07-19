package models

import (
	"github.com/astaxie/beego/orm"
	logpkg "github.com/ckeyer/go-log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/redis"
	"time"
)

var (
	rc     redis.Client
	logfmt = logpkg.MustStringFormatter(
		"%{time:15:04:05} %{shortfile} %{longfunc}%{color} ▶ %{color:reset}[%{color}%{level:.4s}%{color:reset}] %{message}")
	log     = logpkg.MustGetLogger("example")
	db_str  string
	force   = false // force create tables
	verbose = true  // show sql
)

func init() {
	// rc.Addr = beego.AppConfig.String("redis_addr")
	// db_str = beego.AppConfig.String("sql_conn_str")
	db_str = "root:root@/db_blog?charset=utf8"
	rc.Addr = "localhost:6379"
	// backend1 := logpkg.NewLogBackend(os.Stderr, "", 0)
	// backend2 := logpkg.NewLogBackend(os.Stderr, "", 0)

	// // For messages written to backend2 we want to add some additional
	// // information to the output, including the used log level and the name of
	// // the function.
	// backend2Formatter := logpkg.NewBackendFormatter(backend2, logfmt)

	// // Only errors and more severe messages should be sent to backend1
	// backend1Leveled := logpkg.AddModuleLevel(backend1)
	// backend1Leveled.SetLevel(logpkg.ERROR, "")

	// // Set the backends to be used.
	// logpkg.SetBackend(backend1Leveled, backend2Formatter)

}

func RegistDB() {
	orm.Debug = true
	orm.DefaultTimeLoc = time.UTC

	orm.RegisterDataBase("default", "mysql", db_str)
	orm.SetMaxIdleConns("default", 10)
	orm.RegisterModel(new(Blog), new(Tag), new(Review), new(ConnLog), new(BlogTagRelation))

	err := orm.RunSyncdb("default", force, verbose)
	if err != nil {
		panic(err)
	}
}
