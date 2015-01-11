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
		if strings.Count(m.Data, "{") > 0 {
			s += `,"Data":` + m.Data
		} else {
			s += `,"Data":"` + m.Data + `"`
		}
	}
	if "" != m.Desc {
		if strings.Count(m.Desc, "{") > 0 {
			s += `,"Desc":` + m.Desc
		} else {
			s += `,"Desc":"` + m.Desc + `"`
		}
	}
	if 0 != len(m.Supp) {
		for i, v := range m.Supp {
			if 0 == i {
				if strings.Count(v, "{") > 0 {
					s += `,"Supp":[` + v
				} else {
					s += `,"Supp":["` + v
				}
			} else {
				if strings.Count(v, "{") > 0 {
					s += `,` + v
				} else {
					s += `","` + v
				}
			}
		}
		if strings.Count(m.Supp[0], "{") > 0 {
			s += `]`
		} else {
			s += `"]`
		}
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
		if strings.Count(m.Data, "{") > 0 {
			s += `,"Data":` + m.Data
		} else {
			s += `,"Data":"` + base64.URLEncoding.EncodeToString([]byte(m.Data)) + `"`
		}
	}
	if "" != m.Desc {
		if strings.Count(m.Desc, "{") > 0 {
			s += `,"Desc":` + m.Desc
		} else {
			s += `,"Desc":"` + base64.URLEncoding.EncodeToString([]byte(m.Desc)) + `"`
		}
	}
	if 0 != len(m.Supp) {

		for i, v := range m.Supp {
			if 0 == i {
				if strings.Count(v, "{") > 0 {
					s += `,"Supp":[` + v
				} else {
					s += `,"Supp":["` + base64.URLEncoding.EncodeToString([]byte(v))
				}
			} else {
				if strings.Count(v, "{") > 0 {
					s += `,` + v
				} else {
					s += `","` + base64.URLEncoding.EncodeToString([]byte(v))
				}
			}
		}
		if strings.Count(m.Supp[0], "{") > 0 {
			s += `]`
		} else {
			s += `"]`
		}
	}
	s += "}"
	s = strings.Replace(s, "-", "+", -1)
	return
}
