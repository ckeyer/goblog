package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Profile struct {
	Id   int64
	User *User `orm:"reverse(one)"`

	Created int64
	Updated int64
}

// 插入数据库
func (p *Profile) Insert() error {
	o := orm.NewOrm()
	p.Created = time.Now().Unix()
	p.Updated = time.Now().Unix()

	id, err := o.Insert(p)
	if err == nil {
		p.Id = id
	}
	return err
}
