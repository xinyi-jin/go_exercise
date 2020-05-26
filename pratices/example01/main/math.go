package main

import "fmt"

func main() {
	// susu(101, 200)

	// shuixianhua(100, 999)

	jiecheng(3)

	// factorial(3)

}

// 素数
func susu(a int, b int) {
	var slc []int
	// 被除数
	for i := a; i <= b; i++ {
		// 除数
		for j := 2; j < b; j++ {

			// 被除数与除数的大小（被除数需要大于除数）
			if j < i {
				// fmt.Println("a<b")
				if i%j == 0 {
					fmt.Println("this math is not susu.")
					fmt.Printf("%d * %d = %d\n", i/j, j, i)
					slc = append(slc, i)
				}
			}

			if j > i {
				// fmt.Println("a>b")
				if j%i == 0 {
					fmt.Println("this math is not susu.")
					fmt.Printf("%d * %d = %d\n", j/i, i, j)
					slc = append(slc, j)
				}
			}
		}
	}

	slc = RemoveRepByMap(slc)
	fmt.Println(len(slc))
	for _, v := range slc {

		fmt.Printf("%d,", v)

	}
}

// func RemoveRepByMap(slc []int) []int{

// 	result := []int{}
// 	map := map[int]int{}

// 	for _,v := range slc {
// 		l := len(map)
// 		map[v] = 0

// 		if len(map)!=l{
// 			result =  append(result, v)
// 		}
// 	}

// 	return result
// }

// 去重
func RemoveRepByMap(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// 水仙花数

func shuixianhua(a int, b int) {
	var (
		geWei  int
		shiWei int
		baiWei int
		total  int
	)
	// if (a > 100 && b < 1000) {
	// 	new error.Error("参数错误")
	// }
	for i := a; i <= b; i++ {

		geWei = i % 10
		shiWei = (i / 10) % 10
		baiWei = (i / 100) % 10

		total = geWei*geWei*geWei + shiWei*shiWei*shiWei + baiWei*baiWei*baiWei

		if total == i {
			fmt.Printf("%d是水仙花数\n", i)
		}
	}
}

// 阶乘和
func jiecheng(n int) {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += factorial(i)
	}

	fmt.Printf("%d的阶乘和是：%d", n, sum)
}

// 踩坑：当返回值为只有一个的时候，如果要写变量名，必须使用大括号括起来
func factorial(a int) (b int) {
	if a == 0 {
		// fmt.Errorf("error", "参数错误")
		return 1
	}

	recurse := factorial(a - 1)
	resulet := recurse * a
	return resulet
}
