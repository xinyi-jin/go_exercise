package moudle

import (
	"fmt"
	"testing"
)

func TestForRange(t *testing.T) {
	fmt.Println(returnValues())
	fmt.Println(namedReturnValues())
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
}
