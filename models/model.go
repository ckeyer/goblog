package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/goblog/conf"
	"github.com/ckeyer/goblog/lib/logging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vmihailenco/redis"
)

var (
	config  = conf.GetConf()
	rc      *redis.Client
	log     = logging.GetLogger()
	db_str  string
	force   = false // force create tables
	verbose = true  // show sql
)

func init() {
	db_str = config.Mysql.GetConnStr() //"root:root@tcp(d.local:3306)/db_blog?charset=utf8"

	if rc == nil {
		rc = redis.NewTCPClient("d.local:6379", config.Redis.Password, config.Redis.Database)
		// rc = redis.NewTCPClient(config.Redis.GetConnStr(), config.Redis.Password, config.Redis.Database)
		if err := rc.Ping().Err(); err != nil {
			log.Warningf("Redis Ping Failed by %s, ConnStr is: %s\n", err, config.Redis.GetConnStr())
		}
	}
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
