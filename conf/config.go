package conf

import (
	"io/ioutil"
	"os"

	"encoding/json"
	"fmt"
)

type Config struct {
	App     *App     `json:"app"`
	Redis   *Redis   `json:"redis"`
	Mysql   *Mysql   `json:"mysql"`
	WebSite *WebSite `json:"website"`
	WebHook *WebHook `json:"webhook"`
}

type App struct {
	Port int    `json:"port"`
	Name string `json:"name"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Password string `json:"password"`
	Database int64  `json:"database"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Debug    bool   `json:"debug"`
}

type WebSite struct {
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
	Monitors []*Monitor `json:"monitor"`
}

type Monitor struct {
	Branch string `json:"branch"`
	User   string `json:"user"`
	Action string `json:"action"`
	Script string `json:"script"`
}

var config *Config

// 配置文件路径
var path string = "conf/config.json"
var HOOK_SECRET = os.Getenv("HOOK_SECRET")

// init 配置相关初始化
// func init() {
// 	var err error
// 	config, err = load(path)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// SetFilePath 设置配置文件路径
func SetFilePath(filepath string) error {
	path = filepath
	_, err := load(path)
	return err
}

// (c *Config)Load 加载配置
func load(path string) (*Config, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var c Config
	err = json.Unmarshal(bs, &c)
	return &c, err
}

// GetConfig 获取所有配置及错误信息
func GetConfig() (*Config, error) {
	if config != nil {
		return config, nil
	}
	config, err := load(path)
	return config, err
}

// GetConf 仅仅获取配置信息，有错则panic
func GetConf() *Config {
	c, err := GetConfig()
	if err != nil {
		panic(err)
	}
	return c

}

// (m *Mysql)GetConnStr 获取Mysql的连接字符串
func (m *Mysql) GetConnStr() string {
	if m != nil {
		if m.Host == "" {
			m.Host = "localhost"
		}
		if m.Port == 0 {
			m.Port = 3306
		}
		if m.Username == "" {
			m.Username = "root"
		}
		if m.Password == "" {
			m.Password = "root"
		}
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.Username, m.Password, m.Host, m.Port, m.Database)
	}
	return ""
}

// (r *Redis)GetConnStr 获取Redis连接字符串
func (r *Redis) GetConnStr() string {
	if r != nil {
		if r.Host == "" {
			r.Host = "localhost"
		}
		if r.Port == 0 {
			r.Port = 6379
		}
		return fmt.Sprintf("%s:%d", r.Host, r.Port)
	}
	return ""
}
