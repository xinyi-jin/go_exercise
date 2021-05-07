package alg02

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddListNode(l *ListNode, arr []int) {
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			addHeadListNode(l, arr[i])
			continue
		}
		ln := &ListNode{
			Val: arr[i],
		}
		l.Next = ln
		l = ln
	}
}
func addHeadListNode(l *ListNode, val int) {
	l.Val = val
	l = l.Next
}

func Cover(p *ListNode) {
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}

/* 2. 两数相加
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	tail := &ListNode{}
	temp := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + temp
		sum, temp = sum%10, sum/10
		if head == nil { // 头部插入
			head = &ListNode{
				Val: sum,
			}
			tail = head
		} else { // 尾部插入
			tail.Next = &ListNode{
				Val: sum,
			}
			tail = tail.Next
		}
	}
	// 特殊处理最高位有进位情况
	if temp > 0 {
		tail.Next = &ListNode{
			Val: temp,
		}
	}
	return
}
