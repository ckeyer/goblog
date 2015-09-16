package modules

import (
	"github.com/ckeyer/goblog/libs"

	"path/filepath"
)

var (
	log    = libs.GetLogger()
	MyPool = (Pool)(new(MemPooler))
)

// LoadBlogs 加载md博客
func LoadBlogs(dir string) (err error) {
	log.Infof("加载mv博客 %s", dir)
	fs, err := getMDFiles(dir)
	if err != nil {
		return
	}
	MyPool.AddBlogs(fs...)
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

// ReLoadBlogs ...
func ReLoadBlogs(dir string) (err error) {
	log.Infof("更新mv博客 %s", dir)
	fs, err := getMDFiles(dir)
	if err != nil {
		return
	}
	tmp := (Pool)(new(MemPooler))
	count, e := tmp.AddBlogs(fs...)
	log.Notice("共加载文件: ", count, " 错误：", e)

	MyPool = tmp

	return
}
