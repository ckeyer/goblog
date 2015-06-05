package models

import (
	"container/list"
	"github.com/astaxie/beego/orm"
	"time"
)

type Blog struct {
	Id      int64
	Title   string `orm:"size(32)"`
	Page    int    `orm:"default(0)"`
	Summary string
	Content string
	Type    string `orm:"size(12)"`
	Status  int    `orm:"default(0)"`

	Tags []*Tag `orm:"-"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (this *Blog) AddTagName(tag_name string) {
	tag := &Tag{Name: tag_name}
	this.Tags = append(this.Tags, tag)
}
func GetBlogById(id int64) (b *Blog, err error) {
	o := orm.NewOrm()
	b = &Blog{Id: id}

	err = o.Read(b)
	return
}
func (this *Blog) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
	}
	return err
}
func (this *Blog) WriteToDB() (e error) {
	o := orm.NewOrm()
	e = o.Begin()
	// BEGIN
	e = this.Insert()
	if e == nil {
		for _, v := range this.Tags {
			log.Println(v)
			e = v.Get()
			if e != nil {
				break
			}
			InsertBlogTagRelation(this, v)
		}
	}

	// END
	if e != nil {
		e = o.Rollback()
	} else {
		e = o.Commit()
	}
	return
}
func (this *Blog) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}
func (this *Blog) Delete() error {
	o := orm.NewOrm()
	_, err := delBlogInRelation(this)
	if err != nil {
		return err
	}
	_, err = o.Delete(this)
	return err
}

func (this *Blog) ToMap() (bm map[string]string) {
	return
}
func (this *Blog) ToJSON() (s string) {
	return ""
}

func (this *Blog) Read() {
}
func (this *Blog) ReadBlogByID(id_ string) error {
	return nil
}

func (this *Blog) GetNextBlog() (b *Blog, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(b)
	qs.Filter("blog__id__lt", this.Id)
	qs.Limit(1)
	qs.OrderBy("-blog__id")
	// ORDER BY id ASC, profile.age DESC
	err = qs.One(b)
	return
}
func (this *Blog) GetPreviousBlog() (b *Blog, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(b)
	qs.Filter("blog__id__gt", this.Id)
	qs.Limit(1)
	err = qs.One(b)
	return
}

func (this *Blog) GetBlogList(start, stop int) (bs *list.List) {
	return
}
func (this *Blog) GetBlogs(start, stop int) (bs []*Blog) {
	return
}
func (this *Blog) GetBlogsByTagId(tag_id, start, stop int) (bs []*Blog) {
	return
}
