package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Favorite struct {
	Id           int64
	Title        string
	Page         int    `orm:"default(0)"`
	SourceUrl    string `orm:"default("-")"`
	SourceAuthor string `orm:"default("-")"`
	Summary      string
	Content      string
	Status       int    `orm:"default(0)"`
	Source       string `orm:"null "`

	Tags []*Tag  `orm:"-"`  // `orm:"rel(m2m)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func NewFavorite() *Favorite {
	return &Favorite{}
}

func (this *Favorite) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
	}
	return err
}
func (this *Favorite) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}
func (this *Favorite) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
