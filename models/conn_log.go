package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

type ConnLog struct {
	Id       int64
	Domain   string    `orm:"size(16)"`
	Host     string    `orm:"size(16)"`
	Uri      string    `orm:"size(32)"`
	Ip       string    `orm:"size(15)"`
	IpInt    uint32    `orm:"default(0)"`
	Scheme   string    `orm:"size(6)"`
	Method   string    `orm:"size(8)"`
	Protocol string    `orm:"size(16)"`
	DateTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func NewConnLog() *ConnLog {
	return &ConnLog{}
}
func (this *ConnLog) Insert() error {
	o := orm.NewOrm()

	this.IpInt = ip2uint(this.Ip)

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
	}
	return err
}
func ip2uint(ip string) uint32 {
	ips := strings.Split(ip, ".")
	var ipcount uint32 = 0
	for i, v := range ips {
		if n, e := strconv.Atoi(v); e == nil {
			ipn := uint32(n)
			ipcount += ipn << uint((3-i)*8)
		}
	}
	return ipcount
}
