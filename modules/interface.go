package modules

import (
	"sort"
)

type IGroup interface {
	sort.Interface
	FindIndex(string)int
	Insert(int,*Blog)
	GetByIndex(int) *GroupItem
}

// RefreshGroup ...
func RefreshGroup(group IGroup, blog *Blog, names ...string)  {
	for _,name:=range names  {
		index := group.FindIndex(name)
		if index >= 0{
			group.Insert(index, blog)
			log.Debugf("插入 %d  %s ",index,name)
		}else{
			log.Error("获取索引 %s 出错,",name)
		}
	}
	sort.Sort(group)
}

func GetGroupItem(group IGroup, name string) *GroupItem{
	index := group.FindIndex(name)
	return group.GetByIndex(index)
} 
