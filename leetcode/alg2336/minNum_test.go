package alg2336

import (
	"fmt"
	"testing"
)

func TestMinNum(t *testing.T) {
	s := Constructor()
	s.AddBack(2)                 // 2 已经在集合中，所以不做任何变更。
	fmt.Println(s.PopSmallest()) // 返回 1 ，因为 1 是最小的整数，并将其从集合中移除。
	fmt.Println(s.PopSmallest()) // 返回 2 ，并将其从集合中移除。
	fmt.Println(s.PopSmallest()) // 返回 3 ，并将其从集合中移除。
	s.AddBack(1)                 // 将 1 添加到该集合中。
	fmt.Println(s.PopSmallest()) // 返回 1 ，因为 1 在上一步中被添加到集合中，
	// 且 1 是最小的整数，并将其从集合中移除。
	fmt.Println(s.PopSmallest()) // 返回 4 ，并将其从集合中移除。
	fmt.Println(s.PopSmallest()) // 返回 5 ，并将其从集合中移除。
}
