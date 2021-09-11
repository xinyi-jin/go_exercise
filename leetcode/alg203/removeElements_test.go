package alg203

import (
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	var head, want *ListNode = new(ListNode), new(ListNode)
	hd := []int{1, 2, 6, 3, 4, 5, 6}
	head.Val = hd[0]
	wt := []int{1, 2, 3, 4, 5}
	want.Val = wt[0]
	insertTail(head, hd[1:])
	insertTail(want, wt[1:])
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Test_removeElements", args{head, 6}, want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElements(tt.args.head, tt.args.val)
			traverse(got)
			traverse(tt.want)
			/* if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			} */
		})
	}
}
