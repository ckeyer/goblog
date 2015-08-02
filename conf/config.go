package conf

type Config struct {
	App *App
}

type App struct {
	Post int
	Name string
}
