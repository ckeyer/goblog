/*
**/

package models

import (
	// "strings"
	"github.com/astaxie/beego/orm"
	"time"
)

type Tag struct {
	Id       int64
	ParentId int    `orm:"default(0)"`
	Name     string `orm:"size(32);unique"`

	Blogs     []*Blog `orm:"-"`
	BlogCount int64   `orm:"-"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func GetTag(tag_name string) *Tag {
	tag := &Tag{Name: tag_name}
	tag.getByName()
	return tag
}

func (this *Tag) Get() (err error) {
	if this.Name != "" {
		err = this.getByName()
	} else {
		err = this.getById()
	}
	if err != nil {
		return
	}
	// this.getBlogs()
	return
}
func (this *Tag) getById() error {
	o := orm.NewOrm()
	return o.Read(this)
}
func (this *Tag) getByName() error {
	o := orm.NewOrm()
	_, id, err := o.ReadOrCreate(this, "name")
	o.ReadOrCreate(this, "Name")
	if err == nil {
		this.Id = id
	}
	return err

}
func (this *Tag) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
func (this *Tag) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}

// func (this *Tag) getBlogs() {
// 	o := orm.NewOrm()
// 	sql := "select blog.* from blog, blog_tag_relation as bt  where blog.id = bt.blog_id and bt.tag_id = ?"
// 	o.Raw(sql, this.Id).QueryRows(&this.Blogs)

// }
func (this *Tag) GetBlogs(start, count int) (bs []*Blog) {
	o := orm.NewOrm()
	// res := make(orm.Params)
	sql := `select blog.* from blog, blog_tag_relation as bt  where blog.id = bt.blog_id and bt.tag_id = ? order by blog.created desc limit ?,?`
	num, err := o.Raw(sql, this.Id, start, start+count).QueryRows(&bs)
	if num == 0 || err != nil {
		log.Printf("Error Getblogs :Get :%d,Error: %v\n", num, err)
	}
	for _, v := range bs {
		v.getTags()
	}
	return
}

func (this *Tag) GetBlogCount() int64 {
	if this.Blogs != nil {
		this.BlogCount = int64(len(this.Blogs))
		return this.BlogCount
	}
	o := orm.NewOrm()
	cnt, err := o.QueryTable("blog_tag_relation").Filter("tag_id", this.Id).Count()
	if err != nil {
		return 0
	}
	this.BlogCount = cnt
	return cnt

}
func GetHotTags(max int) (tags []*Tag) {
	o := orm.NewOrm()
	sql := "select tag.* from tag,blog_tag_relation as bt  where tag.id = bt.tag_id group by bt.tag_id  order by count(bt.blog_id) desc limit 0,?"
	o.Raw(sql, max).QueryRows(&tags)
	for _, v := range tags {
		v.GetBlogCount()
	}
	return
}
