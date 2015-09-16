package modules

import (
	"sort"
	"strings"
	"sync"
)

type MemPooler struct {
	sync.Mutex
	Blogs      Blogs
	Tags       *Tags
	Categories *Category
	MonthBlog  *MonthBlog
}

// 	 (m *MemPooler) AddBlogs files为文件路径的数组
func (m *MemPooler) AddBlogs(files ...string) (i int, fail int) {
	for i, v := range files {
		blog, err := NewBlog(v)
		if err != nil {
			fail++
			log.Errorf("加载 %d: %s 失败,%s", i, v, err)
			continue
		}
		if m.Tags == nil {
			m.Tags = new(Tags)
		}
		if m.Categories == nil {
			m.Categories = new(Category)
		}
		if m.MonthBlog == nil {
			m.MonthBlog = new(MonthBlog)
		}

		m.Blogs = append(m.Blogs, blog)
		sort.Sort(m.Blogs)

		RefreshGroup(m.Tags, blog, blog.Tags...)
		RefreshGroup(m.Categories, blog, blog.Category...)
		RefreshGroup(m.MonthBlog, blog, blog.Date[:len(blog.Date)-3])
		log.Noticef("加载 %s 成功", v)
	}
	i++
	return
}

// (m *MemPooler) GetBlogs ...
func (m *MemPooler) GetBlogs(start, num int) []*Blog {
	if start < 0 || num < 0 {
		return nil
	}
	if start >= m.Blogs.Len() {
		return nil
	}
	if start+num >= m.Blogs.Len() {
		return m.Blogs[start:]
	}
	return m.Blogs[start : start+num]
}

// GetConf 获取博客总数
func (m *MemPooler) GetCount() int {
	return m.Blogs.Len()
}

// (m *MemPooler) GetBlog 通过名称获取Blog
func (m *MemPooler) GetBlogByName(name string) *Blog {
	for _, b := range m.Blogs {
		if strings.ToLower(b.Name) == strings.ToLower(name) {
			return b
		}
	}
	return nil
}

func (m *MemPooler) GetTags() *Tags {
	return m.Tags
}
func (m *MemPooler) GetCategories() *Category {
	return m.Categories
}
func (m *MemPooler) GetMonthBlogs() *MonthBlog {
	return m.MonthBlog
}
