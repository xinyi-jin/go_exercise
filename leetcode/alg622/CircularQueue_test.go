package alg622

import (
	"testing"
)

func TestMyCircularQueue(t *testing.T) {
	t.Run("TestMyCircularQueue", func(t *testing.T) {
		circularQueue := Constructor(3)
		param_1 := circularQueue.EnQueue(1) // 返回 true
		param_2 := circularQueue.EnQueue(2) // 返回 true
		param_3 := circularQueue.EnQueue(3) // 返回 true
		param_4 := circularQueue.EnQueue(4) // 返回 false，队列已满
		param_5 := circularQueue.Rear()     // 返回 3
		param_6 := circularQueue.IsFull()   // 返回 true
		param_7 := circularQueue.DeQueue()  // 返回 true
		param_8 := circularQueue.EnQueue(4) // 返回 true
		param_9 := circularQueue.Rear()     // 返回 4
		/* if got := Constructor(tt.args.k); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Constructor() = %v, want %v", got, tt.want)
		} */
		if !param_1 || !param_2 || !param_3 || param_4 || param_5 != 3 ||
			!param_6 || !param_7 || !param_8 || param_9 != 4 {
			t.Errorf("param_1 %v, param_2 %v, param_3 %v, param_4 %v, param_5 %v, param_6 %v, param_7 %v, param_8 %v, param_9 %v",
				param_1, param_2, param_3, param_4, param_5, param_6, param_7, param_8, param_9)
		}
	})
}
