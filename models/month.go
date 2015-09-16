package models

import (
	. "github.com/ckeyer/goblog/modules"
)

func GetAllMonth() (ts []string) {

	mblogs := MyPool.GetMonthBlogs()
	ts = make([]string, mblogs.Len())
	for i, v := range mblogs.GetAll() {
		ts[i] = v.Name
	}

	return
}

func GetBlogsByMonth(name string) (bs Blogs) {
	tags := MyPool.GetMonthBlogs()
	return tags.GetByName(name).Blogs
}
