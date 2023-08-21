package binsearch

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "binary search", args: struct {
			array  []int
			target int
		}{array: []int{1, 3, 4, 6, 2, 9, 5}, target: 9}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
