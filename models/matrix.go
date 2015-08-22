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
	Code int                     `form:"code" json:"code"`
	H    int                     `form:"h" json:"h"`
	W    int                     `form:"w" json:"w"`
	Col  int                     `form:"val" json:"val"`
	Data [MATRIX_H][MATRIX_W]int `form:"matrix" json:"matrix"`
}

var key_matrix = "matrix"

func initMatrixRedis() {
	for i := 0; i < MATRIX_H; i++ {
		for j := 0; j < MATRIX_W; j++ {
			v := rc.HSet(key_matrix, fmt.Sprintf("%d_%d", i, j), (fmt.Sprint((i + j) % 5)))
			if v.Err() != nil {
				log.Warningf("Hset Error %s \n", v.Err().Error())
			}

		}
	}
	log.Info("Redis Init Matrix Success.")
}
func GetAllMatrix() (vals *MatrixArray, err error) {
	if b := rc.Exists(key_matrix); b.Err() != nil {
		log.Error(b.Err().Error())
		return
	} else if !b.Val() {
		initMatrixRedis()
	}
	vals, err = getAllToArray()
	return
}
func getAllToArray() (vals *MatrixArray, err error) {
	vals = &MatrixArray{}
	for i := 0; i < MATRIX_H; i++ {
		for j := 0; j < MATRIX_W; j++ {
			b := rc.HGet(key_matrix, fmt.Sprintf("%d_%d", i, j))
			if b.Err() != nil {
				log.Warningf("%v, %s", key_matrix, b.Err().Error())
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
	v := rc.HGet(key_matrix, fmt.Sprintf("%d_%d", h, w))
	count, _ := strconv.Atoi(fmt.Sprintf("%s", v.Val()))
	v2 := rc.HSet(key_matrix, fmt.Sprintf("%d_%d", h, w), (strconv.Itoa(count + 1)))
	return v2.Val(), v2.Err()
}
func (this *MatrixArray) ToJson() string {
	b, e := json.Marshal(*this)
	if e != nil {
		return ""
	}
	return fmt.Sprintf("%s", b)
}
