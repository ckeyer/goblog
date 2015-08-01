package lib

import (
	"fmt"
	"regexp"
)

var (
	REG_EMAIL        = `^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`
	REG_MOBILE_PHONE = `^(1[3|4|5|8][0-9]\d{8})$`
	REG_IDCARD       = `^(\d{17})([0-9]|X|x)$`
)

// 验证Email地址
func IsEmail(email string) bool {
	b, e := regexp.Match(REG_EMAIL, []byte(email))
	if e != nil {
		fmt.Println(e.Error())
	}
	return b
}

// 验证手机号码
func IsMobildPhone(num string) bool {
	b, e := regexp.Match(REG_MOBILE_PHONE, []byte(num))
	if e != nil {
		fmt.Println(e.Error())
	}
	return b
}

// 验证身份证
func IsIDcard(num string) bool {
	b, e := regexp.Match(REG_IDCARD, []byte(num))
	if e != nil {
		fmt.Println(e.Error())
	}
	return b
}
