package alg160

import (
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	var headA, headB, sameAddress, want *ListNode = new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	ha := []int{4, 1, 8, 4, 5}
	headA.Val = ha[0]
	hb := []int{5, 0, 1, 8, 4, 5}
	headB.Val = hb[0]

	insertTail(sameAddress, ha[2:])
	insertTail(headA, ha[1:2])
	insertTailList(headA, sameAddress)
	insertTailList(headB, sameAddress)
	insertTail(headB, hb[1:3])
	insertTail(want, []int{8})

	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{".", args{headA, headB}, want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// cover(tt.args.headA)
			// t.Logf("want value %v", tt.want.Val)
			// cover(tt.args.headB)
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); got != nil {
				t.Logf("got %v", got.Val)
			}
			/* if got := getIntersectionNode(tt.args.headA, tt.args.headB); got != nil && !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("getIntersectionNode() = %v, want %v", got.Val, tt.want.Val)
			} */
		})
	}
}
