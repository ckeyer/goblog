package msg

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type Message struct {
	Code string   `json:"code" form:"code" `
	Data string   `json:"data" form:"data" `
	Desc string   `json:"desc" form:"desc" `
	Supp []string `json:"supp" form:"supp" `
}

func NewMsg(d ...[]byte) (msg *Message, err error) {
	msg = &Message{}
	if len(n) == 1 {
		err = msg.Decode(d[0])
	} else if len(n) != 0 {
		err = errors.New("Too many args")
	}
	return
}
func (m *Message) Decode(data []byte) error {
	m = &Message{}
	err := json.Unmarshal(data, m)
	return err
}
func (m *Message) Encode() ([]byte, error) {
	return json.Marshal(m)
}
