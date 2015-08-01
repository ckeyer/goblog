package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	U_STATUS_NORMAL int8 = 1
	U_STATUS_NEW    int8 = 2
)

var (
	admin = &User{
		Id:     0,
		Name:   "admin",
		Nick:   "admin",
		Email:  "me@ckeyer.com",
		Status: U_STATUS_NORMAL,
	}
)

type User struct {
	Id    int64
	Name  string `orm:"size(32)"`
	Nick  string `orm:"size(32)"`
	Email string `orm:"size(32)"`

	Status int8 `orm:"default(0)"`

	Created int64
	Updated int64

	Profile *Profile `orm:"null;rel(one);on_delete(set_null)"`
}

// 初始化一个匿名用户
func NewAnonymous() *User {
	return &User{
		Name:    "anonymous",
		Nick:    "anonymous",
		Email:   "anonymous@ckeyer.com",
		Profile: &Profile{},
	}
}

// 初始化一个有用户名和邮箱地址的用户
func NewEmailUser(name, email string) *User {
	return &User{
		Name:    name,
		Email:   email,
		Profile: &Profile{},
	}
}

// 插入数据库
func (u *User) Insert() error {
	o := orm.NewOrm()
	u.Status = U_STATUS_NORMAL
	u.Created = time.Now().Unix()
	u.Updated = time.Now().Unix()

	id, err := o.Insert(u)
	if err == nil {
		u.Id = id
	}
	return err
}

// 通过Email获取User，没有则创建
func GetUser(email string) (user *User) {
	o := orm.NewOrm()
	user = &User{Email: email}
	// 三个返回参数依次为：是否新创建的，对象Id值，错误
	if created, id, err := o.ReadOrCreate(user, "Email"); err == nil {
		if created {
			user.Status = U_STATUS_NEW
			log.Info("New Insert an object. Id:", id)
		} else {
			user.Status = U_STATUS_NORMAL
			log.Info("Get an object. Id:", id)
		}
	} else {
		log.Error(err.Error())
	}
	return
}
