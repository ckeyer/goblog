package conf

import (
	"io/ioutil"
	"os"

	"encoding/json"

	"sync"

	"github.com/ckeyer/goblog/libs"
)

var (
	confpath = ""
	log      = libs.GetLogger()
	config   *Config
)

type Config struct {
	sync.RWMutex
	BlogPath string `json:"blog_path"`
}

// init
func init() {

}

// LoadConf 加载配置文件
func LoadConf(path string) (err error) {
	confpath = path
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Errorf("打开文件 %s 失败 %s...", path, err)
		return
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Errorf("读取文件 %s 失败，%s...", path, err)
		return
	}

	c := new(Config)
	err = json.Unmarshal(bs, c)
	if err != nil {
		log.Errorf("解析文件 %s 失败, %s", path, err)
		return
	}

	if config != nil {
		config.Lock()
		defer config.Unlock()
	}
	config = c

	log.Noticef("加载配置 %s 成功",path)
	return
}

// GetConf ...
func GetConf() *Config {
	if config != nil {
		return config
	}
	if confpath != "" {
		if LoadConf(confpath) == nil {
			return config
		}
	}
	panic("获取配置异常...")
}
