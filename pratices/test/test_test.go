package moudle

import (
	"fmt"
	"testing"
)

func TestForRange(t *testing.T) {
	fmt.Println(returnValues())
	fmt.Println(namedReturnValues())
	arr := []int{1, 2, 3}
	for i, v := range arr {
		v += 10
		arr[i] += 10
	}
	fmt.Printf("for range %v\n", arr)
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Printf("for %v\n", arr)
}

/* 将result赋值给返回值（可以理解成Go自动创建了一个返回值retValue，相当于执行retValue = result）
然后检查是否有defer，如果有则执行
返回刚才创建的返回值（retValue）

作者：simpleapples
链接：https://www.jianshu.com/p/79c029c0bd58
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。 */
func returnValues() int {
	var result int
	defer func() {
		result++
		fmt.Println("returnValues defer")
	}()
	return result
}
func namedReturnValues() (result int) {
	defer func() {
		result++
		fmt.Println("namedReturnValues defer")
	}()
	return result
}

// uint 数据溢出会重置0
func TestUint(t *testing.T) {
	var num uint8
	num = 0xFF
	fmt.Println(num)
	num = 0xFE + 1
	fmt.Println("add after", num)
	fmt.Println("-1/9=", -1/9)
}

// 100
func TestCode(t *testing.T) {
	m := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}
	m['a'] = 3
	fmt.Println(len(m))

	// 我们知晓 c 的 ASCII 码是 99，这道题相当于这样
	/* m := [...]int{
		97: 1,
		98: 2,
		99: 3,
	}
	m[97] = 3
	fmt.Println(len(m)) */
}

func TestSlice(t *testing.T) {
	var a []int
	fmt.Printf("a == nil %v\n", a == nil)
	fmt.Printf("[]int{} == nil %v\n", []int{} == nil)
}
func TestTrimSpace(t *testing.T) {
	arr := []byte("hello world")
	res := TrimSpace(arr)
	fmt.Println("arr", string(arr))
	fmt.Println("res", string(res))
}
func TrimSpace(s []byte) []byte {
	b := []byte{}
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}
