package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	MATRIX_H = 7
	MATRIX_W = 30
)

type MatrixArray [MATRIX_H][MATRIX_W]int

type MatrixUpJson struct {
	Msgcode int                     `form:"msgcode",json:"msgcode"`
	H       int                     `form:"h",json:"h"`
	W       int                     `form:"w",json:"w"`
	Col     int                     `form:"val",json:"val"`
	Data    [MATRIX_H][MATRIX_W]int `form:"matrix",json:"matrix"`
}

var key_matrix = "matrix"

func initMatrixRedis() {
	for i := 0; i < MATRIX_H; i++ {
		for j := 0; j < MATRIX_W; j++ {
			rc.Hset(key_matrix, fmt.Sprintf("%d_%d", i, j), []byte(fmt.Sprint((i+j)%5)))
		}
	}
	log.Println("Redis Init Matrix Success.")
}
func GetAllMatrix() (vals *MatrixArray, err error) {
	var b bool
	if b, err = rc.Exists(key_matrix); err != nil {
		log.Println("error redis", err.Error())
		return
	} else if !b {
		initMatrixRedis()
	}
	vals, err = getAllToArray()
	return
}
func getAllToArray() (vals *MatrixArray, err error) {
	vals = &MatrixArray{}
	for i := 0; i < MATRIX_H; i++ {
		for j := 0; j < MATRIX_W; j++ {
			var b []byte
			b, err = rc.Hget(key_matrix, fmt.Sprintf("%d_%d", i, j))
			if err != nil {
				log.Println("error hget ", key_matrix, err.Error())
				initMatrixRedis()
				return
			}
			v, _ := strconv.Atoi(fmt.Sprintf("%s", b))
			vals[i][j] = v % 5
		}
	}
	return
}
func UpdateMatrix(h, w int) (bool, error) {
	bs, _ := rc.Hget(key_matrix, fmt.Sprintf("%d_%d", h, w))
	count, _ := strconv.Atoi(fmt.Sprintf("%s", bs))
	return rc.Hset(key_matrix, fmt.Sprintf("%d_%d", h, w), []byte(strconv.Itoa(count+1)))
}
func (this *MatrixArray) ToJson() string {
	b, e := json.Marshal(*this)
	if e != nil {
		return ""
	}
	return fmt.Sprintf("%s", b)
}
