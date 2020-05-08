package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	flo := Flo{3.141592654}
	self := newFlo(flo)

	fmt.Println(self.fToI())

	fmt.Println("math函数：",math.Ceil(31.14))
}

//float类型的数字
type Flo struct {
	flo float64
}

func (self *Flo) fToI() int {
	str := strconv.FormatFloat(self.flo, 'f', -1, 64)

	fmt.Println("str", str)

	strArr := strings.Split(str, ".")

	rtn := 0

	iNum, _ := strconv.Atoi(strArr[0])
	fmt.Println(iNum)

	split := strings.Split(strArr[1], "E")

	fmt.Println("number",strArr[1])
	f,_ := strconv.Atoi(split[0])
	fmt.Println(f)
	if f > 0 {
		rtn = iNum + 1
		fmt.Println(iNum + 1)
	}

	return rtn

}

func newFlo(self Flo) *Flo {
	return &Flo{
		self.flo,
	}
}
