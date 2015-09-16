package modules

import (
	"sync"
)

type Pool interface {
	sync.Locker
	AddBlogs(path ...string) (int, int) // 添加Blog， 返回加载总数和失败个数

	GetTags() *Tags
	GetCategories() *Category
	GetMonthBlogs() *MonthBlog

	GetBlogByName(name string) *Blog

	GetBlogs(start, num int) []*Blog
	GetCount() int
}
