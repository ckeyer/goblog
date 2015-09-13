package modules

type Category struct {
	Group
}

func (t *Category)GetTag(name string) *GroupItem{
	i:=t.FindIndex(name)
	return t.GetByIndex(i)
}

// (t *Category)GetAll  ...
func (t *Category)GetAll ()[]*GroupItem  {
	all:=make([]*GroupItem, t.Len())
	for i:=0;i<t.Len();i++ {
		all[i] = t.GetByIndex(i)
	}
	return all
}
