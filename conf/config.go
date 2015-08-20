package conf

import (
	"io/ioutil"
	"os"

	"encoding/json"
	"fmt"
)

type Config struct {
	App   *App   `json:"app"`
	Redis *Redis `json:"redis"`
	Mysql *Mysql `json:"mysql"`
}

type App struct {
	Port int    `json:"port"`
	Name string `json:"name"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var config *Config

// 配置文件路径
var path string = "conf/config.json"

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

// GetConfig 获取所有配置
func GetConfig() (*Config, error) {
	if config != nil {
		return config, nil
	}
	config, err := load(path)
	return config, err
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
