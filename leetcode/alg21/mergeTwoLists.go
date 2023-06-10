package alg21

import "fmt"

/* 21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。



示例 1：


输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]


提示：

两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列 */

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 单独内存直接计算 缺点：内存消耗大，执行时间慢 可使用链表指针性质优化
func mergeTwoListsEx(list1 *ListNode, list2 *ListNode) *ListNode {
	var res *ListNode
	for list1 != nil {
		if list2 == nil {
			break
		}
		for list2 != nil {
			if list1.Val < list2.Val {
				res = insertTail(res, list1.Val)
				list1 = list1.Next

				break
			} else {
				res = insertTail(res, list2.Val)
				list2 = list2.Next

			}
		}
	}
	if list1 == nil {
		for list2 != nil {
			res = insertTail(res, list2.Val)
			list2 = list2.Next
		}
	}
	if list2 == nil {
		for list1 != nil {
			res = insertTail(res, list1.Val)
			list1 = list1.Next
		}
	}
	return res
}

func insertTail(list *ListNode, val int) *ListNode {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	if list == nil {
		list = node
	} else {
		temp := list
		for temp != nil {
			if temp.Next == nil {
				temp.Next = node
				break
			}
			temp = temp.Next
		}
	}
	return list
}

func print(list *ListNode) {
	for list != nil {
		fmt.Printf("%v, ", list.Val)
		list = list.Next
	}
}

// 思路：使用链表直接操作，节省空间，提升性能
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	p := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return res.Next
}
