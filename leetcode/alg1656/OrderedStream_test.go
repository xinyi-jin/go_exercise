package alg1656

import (
	"testing"
)

func TestOrderedStream(t *testing.T) {
	t.Run("TestOrderedStream", func(t *testing.T) {
		os := Constructor(5)
		res1 := os.Insert(3, "ccccc") // 插入 (3, "ccccc")，返回 []
		res2 := os.Insert(1, "aaaaa") // 插入 (1, "aaaaa")，返回 ["aaaaa"]
		res3 := os.Insert(2, "bbbbb") // 插入 (2, "bbbbb")，返回 ["bbbbb", "ccccc"]
		res4 := os.Insert(5, "eeeee") // 插入 (5, "eeeee")，返回 []
		res5 := os.Insert(4, "ddddd") // 插入 (4, "ddddd")，返回 ["ddddd", "eeeee"]

		t.Logf("%v %v %v %v %v", res1, res2, res3, res4, res5)
	})
}
