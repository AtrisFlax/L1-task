package main

import (
	"reflect"
	"task10/temps"
	"testing"
)

func TestUniteTempsInGroupsByTenDegree(t *testing.T) {
	type args struct {
		temps []float64
	}
	tests := []struct {
		name    string
		args    args
		want    map[int][]float64
		wantErr bool
	}{
		{name: "OK Case",
			args: struct {
				temps []float64
			}{
				temps: []float64{-25.0, -27.0, -21.0, 13.0, 19.0, 15.5, 24.5},
			},
			want: map[int][]float64{
				-20: {-25.0, -27.0, -21.0},
				10:  {13.0, 19.0, 15.5},
				20:  {24.5},
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := temps.UniteTempsInGroupsByTenDegree(tt.args.temps)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniteTempsInGroupsByTenDegree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniteTempsInGroupsByTenDegree() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmptyArray(t *testing.T) {
	type args struct {
		temps []float64
	}
	tests := []struct {
		name    string
		args    args
		want    map[int][]float64
		wantErr bool
	}{
		{name: "Empty Input Array Case",
			args: struct {
				temps []float64
			}{
				temps: []float64{},
			},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := temps.UniteTempsInGroupsByTenDegree(tt.args.temps)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniteTempsInGroupsByTenDegree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniteTempsInGroupsByTenDegree() got = %v, want %v", got, tt.want)
			}
		})
	}
}
