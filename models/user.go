package models

import ()

type User struct {
	Id    int64
	Name  string
	Nick  string
	Email string

	Profile *Profile
}

type Profile struct {
	Id int64
}

// 初始化一个匿名用户
func NewAnonymous() *User {
	return &User{
		Name:    "anonymous",
		Nick:    "anonymous",
		Email:   "anonymous@ckeyer.com",
		Profile: &Perfile{},
	}
}

// 初始化一个有用户名和邮箱地址的用户
func NewEmailUser(name, email string) *User {
	return &User{
		Name:    name,
		Email:   email,
		Profile: &Perfile{},
	}
}
