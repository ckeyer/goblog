package models

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type Message struct {
	Code string
	Data string
	Desc string
	Supp []string
}

func NewMsg() (msg *Message) {
	msg = &Message{}
	return
}
func DecodeJson(data []byte) (msg *Message, err error) {
	s := strings.Replace(string(data), "+", "-", -1)
	bs, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return
	}
	msg = &Message{}
	err = json.Unmarshal(bs, msg)
	if nil != err {
		return
	}
	return
}
func (this *Message) checkErr() {
	this.Code = "err"
}
func (this *Message) ToBytes() (b []byte, e error) {
	b, e = json.Marshal(this)
	return
}
func (this *Message) ToString() (s string) {
	if b, e := this.ToBytes(); e != nil {
		return ""
	} else {
		return string(b)
	}
}
func (this *Message) ToBase64String() (s string) {
	b, e := this.ToBytes()
	if e != nil {
		return ""
	}
	s = base64.StdEncoding.EncodeToString(b)
	s = strings.Replace(s, "-", "+", -1)
	return
}
