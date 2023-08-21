package task11

import (
	"reflect"
	"sort"
	"testing"
)

func TestHashIntersection(t *testing.T) {
	type args[T comparable] struct {
		a []T
		b []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{name: "Hash intersection", args: struct {
			a []int
			b []int
		}{
			a: []int{3, 1, 2, 5, 4, 7, 6},
			b: []int{0, 0, 0, 4, 5, 9, 8}},
			want: []int{4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HashIntersection[int](tt.args.a, tt.args.b)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleIntersection(t *testing.T) {
	type args[T comparable] struct {
		a []T
		b []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{name: "Simple intersection", args: struct {
			a []int
			b []int
		}{
			a: []int{3, 1, 2, 5, 4, 7, 6},
			b: []int{0, 0, 0, 4, 5, 9, 8}},
			want: []int{4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimpleIntersection[int](tt.args.a, tt.args.b)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
