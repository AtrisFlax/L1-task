package test

import (
	"reflect"
	squaresSum "task3"
	"testing"
)

func TestStreamsCalculate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2^2 + 4^2 + 6^2 + 8^2 + 10^2 == 220",
			args: struct{ nums []int }{nums: []int{2, 4, 6, 8, 10}},
			want: 220,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squaresSum.CalculatStreams(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
