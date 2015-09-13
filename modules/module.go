package modules

import (
	"github.com/ckeyer/goblog/libs"

	"path/filepath"
)

var (
	log = libs.GetLogger()
	pool = (Pool)(new(MemPooler))
)

// LoadBlogs 加载md博客
func LoadBlogs(dir string)(err error) {
	log.Infof("加载mv博客 %s", dir)
	fs,err:=getMDFiles(dir)
	if err!=nil {
		return 
	}
	pool.AddBlogs(fs...)
	return
}

// GetMDFiles ...
func getMDFiles(dir string) (files []string, err error) {
	files, err = filepath.Glob(dir)
	if err != nil {
		log.Errorf("%s", err)
		return
	}
	return 
}
