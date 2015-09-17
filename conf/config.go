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
	BlogDir string   `json:"blog_dir"`
	WebSite *WebSite `json:"website"`
	WebHook *WebHook `json:"webhook"`
}

type WebSite struct {
	Title        string   `json:"title"`
	Keywords     []string `json:"keywords"`
	Description  string   `json:"description"`
	HostUrl      string   `json:"host_url"`
	FileUrl      string   `json:"file_url"`
	JsUrl        string   `json:"js_url"`
	CssUrl       string   `json:"css_url"`
	ImgUrl       string   `json:"img_url"`
	CustomJsUrl  string   `json:"custom_js_url"`
	CustomCssUrl string   `json:"custom_css_url"`
	CustomImgUrl string   `json:"custom_img_url"`
	EnableDomain []string `json:"enable_domain"`
}
type WebHook struct {
	Repos    string     `json:"repos"`
	Secret   string     `json:"secret"`
	Monitors []*Monitor `json:"monitor"`
}

type Monitor struct {
	Branch string `json:"branch"`
	User   string `json:"user"`
	Action string `json:"action"`
	Script string `json:"script"`
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

	log.Noticef("加载配置 %s 成功", path)
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
	log.Errorf("%s, %#v", confpath, config)
	panic("获取配置异常...")
}
