package task23

import (
	"reflect"
	"testing"
)

func TestRemoveOrdered(t *testing.T) {
	type args struct {
		slice []int
		ind   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ordered remove",
			args: struct {
				slice []int
				ind   int
			}{slice: []int{1, 2, 3, 4}, ind: 2}, want: []int{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveOrdered(tt.args.slice, tt.args.ind); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveOrdered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveUnordered(t *testing.T) {
	type args struct {
		slice []int
		idx   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "unordered remove",
			args: struct {
				slice []int
				idx   int
			}{slice: []int{1, 2, 3, 4}, idx: 2}, want: []int{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveUnordered(tt.args.slice, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveUnordered() = %v, want %v", got, tt.want)
			}
		})
	}
}
