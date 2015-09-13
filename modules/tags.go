package modules

type Tags struct {
	Group
}

func (t *Tags)GetTag(name string) *GroupItem{
	i:=t.FindIndex(name)
	return t.GetByIndex(i)
}

// (t *Tags)GetAll  ...
func (t *Tags)GetAll ()[]*GroupItem  {
	all:=make([]*GroupItem, t.Len())
	for i:=0;i<t.Len();i++ {
		all[i] = t.GetByIndex(i)
	}
	return all
}
