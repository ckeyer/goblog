package models

import (
// "encoding/json"
)

type jsonMatrix struct {
	X int `json:"x"`
}

func GetAllMatrix() [][]int {

	if b, e := rc.Exists("matrix"); b && e == nil {

	} else {

	}
	return nil
}
