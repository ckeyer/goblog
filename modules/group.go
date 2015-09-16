/* group by
 */

package modules

import (
	"sort"
	"strings"
)

type GroupItem struct {
	weighting int64 // 权重，用于组内排序
	Name      string
	Blogs     Blogs
}

func (g *GroupItem) Count() int {
	return len(g.Blogs)
}
func (g *GroupItem) GetName() string {
	return g.Name
}

type Group []*GroupItem

// (g *Groups)Len ...
func (g Group) Len() (count int) {
	return len(g)
}

// (g *Groups)Less ...
func (g Group) Less(i, j int) bool {
	if g[i].weighting > g[j].weighting {
		return true
	}
	return false
}

// (g *Group)Sort ...
func (g Group) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

// (g *Group)FindAll  查找出所有标签的索引
func (g *Group) FindIndex(name string) int {
	for i, v := range *g {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			return i
		}
	}
	gi := new(GroupItem)
	gi.Name = name
	*g = append(*g, gi)
	return len(*g) - 1
}

// (g *Group)AddItem ...
func (g *Group) AddBlog(index int, weitht int64, b *Blog) {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorf("插入标签时出错 %s", err)
		}

	}()
	gi := (*g)[index]
	gi.weighting = weitht
	gi.Blogs = append(gi.Blogs, b)
	sort.Sort(gi.Blogs)
}

func (g Group) Sort() {
	sort.Sort(g)
}

// GetByIndex ...
func (g Group) GetByIndex(i int) *GroupItem {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorf("插入标签时出错 %s", err)
		}
	}()
	return g[i]
}

// (g *Group)GetByName ...
func (g Group) GetByName(name string) *GroupItem {
	index := g.FindIndex(name)
	return g.GetByIndex(index)
}

// (t *Tags)AddBlogs 更新标签数据
func (g *Group) Insert(index int, b *Blog) {
	weitht := (int64)(g.GetByIndex(index).Count()) + 1
	g.AddBlog(index, weitht, b)
}
