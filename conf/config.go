package conf

import (
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
