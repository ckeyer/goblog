package modules

type MemPooler struct {
	Blogs []*Blog
	Tags *Tags
	Categories *Category
	MonthBlog *MonthBlog
}

// 	 (m *MemPooler) AddBlogs files为文件路径的数组 
func (m *MemPooler) AddBlogs(files ...string) (i int, fail int ) {
	for i, v := range files {
		blog,err := NewBlog(v)
		if err != nil {
			fail ++
			log.Errorf("加载 %d: %s 失败,%s",i, v,err)
			continue
		}
		if m.Tags == nil{
			m.Tags = new(Tags)
		}
		if m.Categories == nil{
			m.Categories = new(Category)
		}
		if m.MonthBlog == nil{
			m.MonthBlog = new(MonthBlog)
		}
		
		RefreshGroup(m.Tags, blog, blog.Tags...)
		RefreshGroup(m.Categories, blog, blog.Category...)
		RefreshGroup(m.MonthBlog, blog, blog.Date)
		log.Noticef("加载 %s 成功",v)
	}
	i++
	return
}

// (m *MemPooler) GetBlogs ...
func (m *MemPooler) GetBlogs(start, num int) []*Blog {
	return nil
}

func (m *MemPooler) GetTags() *Tags{
	return m.Tags
}
func (m *MemPooler) GetCategory() *Category{
	return m.Categories
}
func (m *MemPooler) GetMonthBlog() *MonthBlog{
	return m.MonthBlog
}
