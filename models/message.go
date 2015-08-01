package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id   int64
	From *User  `orm:"rel(fk)"`
	To   string `orm:"size(16)"`

	Content  string `orm:"size(256)"`
	Password string `orm:"size(32)"`
	Created  int64
	Updated  int64
}

// 插入数据库
func (m *Message) Insert() error {
	o := orm.NewOrm()

	m.Created = time.Now().Unix()
	m.Updated = time.Now().Unix()

	id, err := o.Insert(m)
	if err == nil {
		m.Id = id
	}
	return err
}
