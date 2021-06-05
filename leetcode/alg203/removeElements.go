package alg203

import "fmt"

/* 203. 移除链表元素
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-decreasing-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：循环迭代找出符合的节点，直接删除 思路没问题，删除节点的时候要保留删除节点前的链表数据
// 解决问题的关键在于三个节点 前 中 后 使当前节点一直处于中节点位置，可直接操作前后节点吃掉中节点
func removeElements(head *ListNode, val int) *ListNode {
	prev := &ListNode{Next: head}
	for temp := prev; temp.Next != nil; {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return prev.Next
}

// 尾部插入 数组数据
func insertTail(l *ListNode, arr []int) {
	var tail = l
	for _, v := range arr {
		list := &ListNode{Val: v}
		tail.Next = list
		tail = list
	}
}

func traverse(p *ListNode) {
	for p != nil {
		fmt.Printf("%v ", p.Val)
		p = p.Next
	}
	fmt.Println()
}
