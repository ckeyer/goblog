package libs

import logpkg "github.com/ckeyer/go-log"

var (
	log *logpkg.Logger
)

func init() {
	if log == nil {
		log = logpkg.GetDefaultLogger("blog")
	}
}

func GetLogger() *logpkg.Logger {
	return log
}
