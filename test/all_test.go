package test

import (
	"fmt"
	"reflect"
	"testing"
)

type Reflect struct {
	name    string
	age     int
	address string
}

func TestPKGReflect(t *testing.T) {
	var roll interface{}
	r := &Reflect{"name", 24, "zhengzhou,henan"}
	roll = r
	nt := reflect.TypeOf(roll)
	nv := reflect.ValueOf(roll)
	fmt.Printf("type %v \nvalue %v \n", nt, nv)
	res := roll.(*Reflect)
	fmt.Printf("res %v \n", res)

	// float64'/'运算后转int,默认舍弃小数位（向下取整）
	var f1, f2, f3 float64 = 1.8, 1.3, 1.5
	fmt.Printf("f1:%v, f2:%v, f3:%v \n", int(f1), int(f2), int(f3))
}
