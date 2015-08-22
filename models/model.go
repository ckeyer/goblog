package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/goblog/conf"
	"github.com/ckeyer/goblog/lib/logging"
	//	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	redis "gopkg.in/redis.v3"
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

	if rc == nil || rc.Ping().Err() != nil {
		rc = redis.NewClient(&redis.Options{
			Addr:     config.Redis.GetConnStr(),
			Password: config.Redis.Password,
			DB:       config.Redis.Database,
		})
		// rc = redis.NewTCPClient(config.Redis.GetConnStr(), config.Redis.Password, config.Redis.Database)
		if err := rc.Ping().Err(); err != nil {
			log.Warningf("Redis Ping Failed by %s, ConnStr is: %s, %s ,%d\n", err, config.Redis.GetConnStr(), config.Redis.Password, config.Redis.Database)
		} else {
			log.Notice("Redis Connect Success...")
		}
	}
}

func RegistDB() {
	orm.Debug = config.Mysql.Debug
	orm.DefaultTimeLoc = time.UTC

	orm.RegisterDataBase("default", "mysql", db_str)
	orm.SetMaxIdleConns("default", 10)
	orm.RegisterModel(new(Blog), new(Tag), new(Review),
		new(ConnLog), new(BlogTagRelation), new(User),
		new(Profile), new(Message))

	err := orm.RunSyncdb("default", force, verbose)
	if err != nil {
		log.Panic(err)
	} else {
		log.Notice("Mysql Orm Build Success ")
	}
}
