package models

import (
	. "github.com/ckeyer/goblog/modules"
)

// GetBlogs
func GetBlogs(start, count int) (bs []*Blog) {
	bs = MyPool.GetBlogs(start, count)
	return
}

// GetBlog
func GetBlog(name string) *Blog {
	b := MyPool.GetBlogByName(name)
	return b
}

// GetCount
func GetCount() int {
	return MyPool.GetCount()
}
