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

// 自定义测试
func TestCustom(t *testing.T) {
	// 测试switch 分支匹配后，还会不会继续匹配
	var cd = 30
	switch {
	case cd < 30:
		fmt.Println("cd<30")
	case cd < 60:
		fmt.Println("cd<60")
		return // 测试switch 分支匹配后，return后，还会不会执行外层逻辑
	default:
		fmt.Println("cd<???")
	}
	fmt.Println("ending....")
}
