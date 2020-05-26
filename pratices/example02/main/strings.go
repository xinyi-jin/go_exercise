package main

import (
	"fmt"
	"math"
)

func main() {

	// jiujiuCFB()
	// isWanshu(6)

	// slc := testWanshu(1, 1000)
	// fmt.Print("一千以内的完数：")
	// for _, v := range slc {
	// 	fmt.Printf("%d,", v)
	// }

	// flag := isHuiWen("12345689k954321")
	// if flag {
	// 	fmt.Println("huiwenshu")
	// }
	// if !flag {
	// 	fmt.Println("not huiwenshu ")
	// }

	// stringsCount("12345 6ahsdjAHD FJAHK!#$$ ^$#^&%$&")

	add("100", "12111")
}

// 九九乘法表
func jiujiuCFB() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func testWanshu(a int, b int) (all []int) {
	var flag bool

	for i := a; i <= b; i++ {
		flag = isWanshu(i)
		if flag {
			all = append(all, i)
		}
	}
	return
}

// 判断是否是完数
func isWanshu(n int) (flag bool) {
	// 求因数和
	sum := yinshuSum(n)
	// fmt.Println("源数字：", n)
	// fmt.Println("因数之和：", sum)
	if sum == n {
		// fmt.Println("这是一个完数")
		flag = true
	}
	return
}

// 因数之和----完数的因数，算1而不算他本身数字
func yinshuSum(n int) (sum int) {

	// 先计算除了一和他本身的因数和
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			sum = sum + i + (n / i)
			// fmt.Println("因数：", i, (n / i))
		}
	}
	sum += 1
	return
}

// 回文数
func isHuiWen(str string) (flag bool) {
	flag = true

	for i := 0; i < len(str)/2; i++ {
		for j := len(str) - 1 - i; j > len(str)/2; j-- {
			// fmt.Println(str[i], str[j])
			if str[i] == str[j] {
				// fmt.Println("huiwen ")
				break
			} else {
				// fmt.Println("not huiwen")
				flag = false
				return
			}
		}
	}

	return
}

func stringsCount(str string) {
	var (
		spaceCount       int //空格数
		numberCount      int //数字个数
		zimuCount        int //字母个数
		specialCharCount int //特殊字符个数

	)
	// fmt.Println("请输入需要统计的内容：")
	// fmt.Scanf("%s", &str)

	for _, v := range str {
		// fmt.Printf("%d%T\n", v, v)
		switch {
		case v == ' ':
			// fmt.Println("空格")
			spaceCount++
		case v >= '0' && v <= '9':
			// fmt.Println("数字")
			numberCount++
		case v >= 'a' && v <= 'z':
			// fmt.Println("小写字母")
			zimuCount++
		case v >= 'A' && v <= 'Z':
			// fmt.Println("大写字母")
			zimuCount++
		default:
			// fmt.Println("特殊字符")
			specialCharCount++
		}
	}
	fmt.Printf("空格:%d, 数字:%d, 字母:%d, 特殊字符:%d", spaceCount, numberCount, zimuCount, specialCharCount)

}

type subInt string

var (
	ints1     []int32
	ints2     []int32
	solNum    int
	jinWeiNum int
	sum       int

	result []int
)

func add(a subInt, b subInt) {
	for _, i := range a {
		ints1 = append(ints1, i)
	}
	for _, j := range b {
		ints2 = append(ints2, j)
	}

	if len(ints1) <= len(ints2) {
		for i := 0; i < len(ints1)-1; i++ {
			solNum = int(ints1[i]) + int(ints2[i]) + jinWeiNum
			if solNum >= 10 {
				jinWeiNum = solNum / 10
			} else {
				jinWeiNum = 0
			}

			result = append(result, int(solNum))
		}
	}

	fmt.Println(result)
}
