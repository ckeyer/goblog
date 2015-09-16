package modules

import (
	"encoding/json"
)

type Header struct {
	Name     string   `json:"name"`
	Author   string   `json:"author"`
	Head     string   `json:"head"`
	Date     string   `json:"date"`
	Title    string   `json:"title"`
	Tags     []string `json:"tags"`
	Category []string `json:"category"`
	Status   string   `json:"status"`
	Summary  string   `json:"summary"`
}

// (h *Header) 获取blog头部信息
func (h *Header) Load(bs []byte) (err error) {
	err = json.Unmarshal(bs, h)
	if err != nil {
		log.Errorf("头部信息解析错误  %s", err)
	}
	return
}
