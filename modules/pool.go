package modules

type Pool interface {
	AddBlogs(path ...string) (int, int) // 添加Blog， 返回加载总数和失败个数
	GetTags()(*Tags)
	GetCategory() (*Category)
	GetMonthBlog() *MonthBlog
	GetBlogs(start,num int) []*Blog
}
