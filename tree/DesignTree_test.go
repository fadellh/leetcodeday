package tree

import (
	"testing"
)

func TestTree_Insert(t *testing.T) {
	type fields struct {
		root *Node
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				root: &Node{
					val: 7,
					leftChild: &Node{
						val: 5,
						leftChild: &Node{
							val:        3,
							leftChild:  nil,
							rightChild: nil,
						},
						rightChild: &Node{
							val:        4,
							leftChild:  nil,
							rightChild: nil,
						},
					},
					rightChild: nil,
				},
			},
			args: args{
				val: 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Tree{
				root: tt.fields.root,
			}
			this.Insert(tt.args.val)
		})
	}
}

func TestImplemen(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "HSa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Implemen()
		})
	}
}
