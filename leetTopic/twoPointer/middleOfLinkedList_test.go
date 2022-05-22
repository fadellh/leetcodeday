package twopointer

import (
	"reflect"
	"testing"
)

func TestMiddleOfLinkedList(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: &ListNode{},
					},
				},
			},
		},
	}

	want := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  5,
			Next: &ListNode{},
		},
	}

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test case 1",
			args: args{list},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiddleOfLinkedList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MiddleOfLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddleNode(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: &ListNode{},
					},
				},
			},
		},
	}

	want := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  5,
			Next: &ListNode{},
		},
	}

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test case 1",
			args: args{list},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiddleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
