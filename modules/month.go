package modules

import (
	"time"
)
 
type MonthBlog struct {
	Group
}

// (t *MonthBlog)AddBlogs 更新标签数据
func (m *MonthBlog)Insert(index int  ,b *Blog)  {
	var weitht int64
	t,err:=time.Parse( "2006-01-02", b.Date)
	if err != nil {
		log.Errorf("格式化时间 %s 错误 %s",b.Date,err)
		return
	}
	weitht = t.Unix()
	log.Debug("Insert Time ",weitht," ",b.Date," ",t.String())
	m.AddBlog(index ,weitht , b)
}

func (t *MonthBlog)GetTag(name string) *GroupItem{
	i:=t.FindIndex(name)
	return t.GetByIndex(i)
}

// (t *MonthBlog)GetAll  ...
func (t *MonthBlog)GetAll ()[]*GroupItem  {
	all:=make([]*GroupItem, t.Len())
	for i:=0;i<t.Len();i++ {
		all[i] = t.GetByIndex(i)
	}
	return all
}
