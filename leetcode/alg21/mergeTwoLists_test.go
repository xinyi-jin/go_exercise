package alg21

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Test_mergeTwoLists_01", args{
			list1: &ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						4,
						nil,
					},
				},
			},
			list2: &ListNode{
				1,
				&ListNode{
					3,
					&ListNode{
						4,
						nil,
					},
				},
			},
		}, &ListNode{
			1,
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							&ListNode{
								4,
								nil,
							},
						},
					},
				},
			},
		}},
		{"Test_mergeTwoLists_02", args{
			list1: nil,
			list2: nil,
		}, nil},
		{"Test_mergeTwoLists_03", args{
			list1: nil,
			list2: &ListNode{
				0,
				nil,
			},
		}, &ListNode{
			0,
			nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			print(tt.args.list1)
			fmt.Println()
			print(tt.args.list2)
			fmt.Println()
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				print(got)
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_print(t *testing.T) {
	type args struct {
		list *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test_print", args{list: &ListNode{
			1,
			&ListNode{
				2,
				&ListNode{
					4,
					nil,
				},
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			print(tt.args.list)
		})
	}
}
