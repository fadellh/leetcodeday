package linkedlist

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		arr []int
		pos int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				arr: []int{4, 3, 2, 1},
				pos: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Prepare LinkedList
			first := initList()
			cy1 := &ListNode{}
			cy2 := &ListNode{}
			for i, v := range tt.args.arr {
				if i == tt.args.pos {
					cy1 = first.AddFront(v)
					continue
				}
				//Create cycle
				if i == len(tt.args.arr)-1 {
					cy2 = first.AddFront(v)
					continue
				}
				first.AddFront(v)
			}

			cy1.Next = cy2

			if got := hasCycle(first.Head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
