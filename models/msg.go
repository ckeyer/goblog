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
func DecodeJson(data string) (msg *Message) {
	msg = &Message{}
	data = strings.Replace(data, "+", "-", -1)
	// func Replace(s, old, new string, n int) string
	err := json.Unmarshal([]byte(data), msg)
	if nil != err {
		msg.checkErr()
	}
	var bysTmp []byte
	bysTmp, _ = base64.URLEncoding.DecodeString(msg.Code)
	msg.Code = string(bysTmp)
	bysTmp, _ = base64.URLEncoding.DecodeString(msg.Data)
	msg.Data = string(bysTmp)
	bysTmp, _ = base64.URLEncoding.DecodeString(msg.Desc)
	msg.Desc = string(bysTmp)
	for i, v := range msg.Supp {
		bysTmp, _ = base64.URLEncoding.DecodeString(v)
		msg.Supp[i] = string(bysTmp)
	}
	return
}
func (m *Message) checkErr() {
	m.Code = "err"
}
func (m *Message) ToString() (s string) {
	if "" != m.Code {
		s = `{"Code":"` + m.Code + `"`
	} else {
		s = `{"Code":"bmls"`
	}
	if "" != m.Data {
		s += `,"Data":"` + m.Data + `"`
	}
	if "" != m.Desc {
		s += `,"Desc":"` + m.Desc + `"`
	}
	if 0 != len(m.Supp) {
		for i, v := range m.Supp {
			if 0 == i {
				s += `,"Supp":["` + v
			} else {
				s += `","` + v
			}
		}
		s += `"]`
	}
	s += "}"
	return
}
func (m *Message) ToBase64String() (s string) {
	if "" != m.Code {
		s = `{"Code":"` + base64.URLEncoding.EncodeToString([]byte(m.Code)) + `"`
	} else {
		s = `{"Code":"bmls"`
	}
	if "" != m.Data {
		s += `,"Data":"` + base64.URLEncoding.EncodeToString([]byte(m.Data)) + `"`
	}
	if "" != m.Desc {
		s += `,"Desc":"` + base64.URLEncoding.EncodeToString([]byte(m.Desc)) + `"`
	}
	if 0 != len(m.Supp) {
		for i, v := range m.Supp {
			if 0 == i {
				s += `,"Supp":["` + base64.URLEncoding.EncodeToString([]byte(v))
			} else {
				s += `","` + base64.URLEncoding.EncodeToString([]byte(v))
			}
		}
		s += `"]`
	}
	s += "}"
	s = strings.Replace(s, "-", "+", -1)
	return
}
