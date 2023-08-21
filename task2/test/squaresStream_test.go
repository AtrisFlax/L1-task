package test

import (
	"reflect"
	squares "task2"
	"testing"
)

func TestStreamsCalculate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "2, 4, 6, 8, 10",
			args: struct{ nums []int }{nums: []int{2, 4, 6, 8, 10}},
			want: []int{4, 16, 36, 64, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squares.CalculateStreams(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
