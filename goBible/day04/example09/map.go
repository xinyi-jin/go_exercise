package main

import "fmt"

func main() {
	a := make(map[string]map[string]string)
	a["001"] = make(map[string]string)
	a["001"]["name"] = "jxy"

	traverse(a)
	modifyMap(a, "001", "name", "zhuhe")

	traverse(a)

	testSliceMap()

}

/* 遍历输出map */
func traverse(m map[string]map[string]string) {
	for k, v := range m {
		fmt.Println(k)
		for k1, v1 := range v {
			fmt.Println("key:", k1, "value:", v1)
		}
	}
}

/* 修改map元素的值 */
func modifyMap(m map[string]map[string]string, s string, str string, str2 string) {
	_, ok := m[s]
	if !ok {
		m[s] = make(map[string]string)
	}

	m[s][str] = str2
	return
}

func testSliceMap() {
	var a []map[int]int
	a = make([]map[int]int, 5)
	if a[0] == nil {
		a[0] = make(map[int]int)
	}
	a[0][10] = 100

	fmt.Println(a[0][10])
}
