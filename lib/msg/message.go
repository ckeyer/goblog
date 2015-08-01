package msg

import (
	"encoding/json"
	"errors"
)

type Message struct {
	Code string   `json:"code" form:"code" `
	Data string   `json:"data" form:"data" `
	Desc string   `json:"desc" form:"desc" `
	Supp []string `json:"supp" form:"supp" `
}

func NewMsg(d ...[]byte) (msg *Message, err error) {
	msg = &Message{}
	if len(d) == 1 {
		err = msg.Decode(d[0])
	} else if len(d) != 0 {
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
func (m *Message) String() (s string, err error) {
	bs, err := m.Bytes()
	s = string(bs)
	return
}
func (m *Message) Bytes() (bs []byte, err error) {
	bs, err = m.Encode()
	return
}
