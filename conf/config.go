package conf

import (
	"io/ioutil"
	"os"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/ckeyer/go-lib"
)

type Config struct {
	App   *App
	Redis *lib.RedisConfig
	Mysql *lib.MySqlConfig
}

type App struct {
	Post int
	Name string
}

var config *Config

// init 配置相关初始化
func init() {
	path := beego.AppConfig.String("config_path")
	if path == "" {
		return
	}
	config = new(Config)
	config.Load(path)
}

// (c *Config)Load 加载配置
func (c *Config) Load(path string) (err error) {
	var conf Config

	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	err = json.Marshal(bs, &conf)
	c = &conf
	return
}
