package alg641

import (
	"testing"
)

func TestMyCircularDeque(t *testing.T) {
	t.Run("TestMyCircularDeque", func(t *testing.T) {
		circularDeque := Constructor(3)         // 设置容量大小为3
		param_1 := circularDeque.InsertLast(1)  // 返回 true
		param_2 := circularDeque.InsertLast(2)  // 返回 true
		param_3 := circularDeque.InsertFront(3) // 返回 true
		param_4 := circularDeque.InsertFront(4) // 已经满了，返回 false
		param_5 := circularDeque.GetRear()      // 返回 2
		param_6 := circularDeque.IsFull()       // 返回 true
		param_7 := circularDeque.DeleteLast()   // 返回 true
		param_8 := circularDeque.InsertFront(4) // 返回 true
		param_9 := circularDeque.GetFront()     // 返回 4

		if !param_1 || !param_2 || !param_3 || param_4 || param_5 != 2 ||
			!param_6 || !param_7 || !param_8 || param_9 != 4 {
			t.Errorf("param_1 %v, param_2 %v, param_3 %v, param_4 %v, param_5 %v, param_6 %v, param_7 %v, param_8 %v, param_9 %v",
				param_1, param_2, param_3, param_4, param_5, param_6, param_7, param_8, param_9)
		}
	})
}
