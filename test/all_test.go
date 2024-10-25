package test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
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
	/* {
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
	} */

	// 测试时间戳 （根据秒级时间戳，计算当日点时刻）
	{
		ts := time.Now().Unix()
		regTime := time.Unix(ts, 0).Local()
		regTime = time.Date(regTime.Year(), regTime.Month(), regTime.Day()+1, 0, 0, 0, int(0), regTime.Location())

		fmt.Println("zore time:", regTime.Format("2006-01-02 15:04:05"))
	}

	// 身份证号计算对应的出生年和性别
	{
		idNumber := "410225199705151518"
		var tmp string
		for i := 6; i < 10; i++ {
			tmp += string(idNumber[i])
		}
		println("tmp", tmp)
		birthYear, err := strconv.Atoi(tmp)
		if err != nil {
			fmt.Errorf("Atoi err %v", err)
		}
		sex, err := strconv.Atoi(string(idNumber[16]))
		if err != nil {
			fmt.Errorf("Atoi err %v", err)
		}
		println("birthYear", fmt.Sprintf("%d", birthYear), "sex", sex)

		println(time.Unix(0, 0).Format("2006-01-02 15:04:05"))
	}
}
