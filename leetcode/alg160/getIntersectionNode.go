package alg160

import "fmt"

/* 160. 相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-decreasing-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//  思路：Hash实现 (未完全实现) {根据结点中存储的值判定，会存在特殊情况，导致结果不正确；可根据当前结点指针地址判定，此值是唯一值}
func getIntersectionNodeNotFinish(headA, headB *ListNode) *ListNode {
	aMap := map[int]*ListNode{}
	for {
		if headA != nil {
			aMap[headA.Val] = headA
		}
		if headA.Next != nil {
			headA = headA.Next
			continue
		} else {
			break
		}
	}
	for {
		if headB != nil {
			if _, ok := aMap[headB.Val]; ok {
				return headB
			}
		}
		if headB.Next != nil {
			headB = headB.Next
			continue
		} else {
			break
		}
	}
	return nil
}

// 官方题解 Hash
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	aMap := map[*ListNode]bool{}
	for temp := headA; temp != nil; temp = temp.Next {
		aMap[temp] = true
	}
	for temp := headB; temp != nil; temp = temp.Next {
		if aMap[temp] {
			return temp
		}
	}
	return nil
}

// 官方题解 双指针
func getIntersectionNodeEx(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
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

// 尾部插入 链表节点数据
func insertTailList(l, in *ListNode) {
	var tail = l
	for temp := tail; ; temp = temp.Next {
		if temp == nil {
			tail = &ListNode{}
			break
		}
	}
	for temp := in; temp != nil; temp = temp.Next {
		tail.Next = temp
		tail = &ListNode{}
	}
}

func cover(p *ListNode) {
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
