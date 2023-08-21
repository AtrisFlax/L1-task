package task24

import "testing"

func TestPoint_Distance(t *testing.T) {
	type fields struct {
		x float64
		y float64
	}
	type args struct {
		p2 *Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{name: "distance calculate",
			fields: struct {
				x float64
				y float64
			}{x: 0, y: -3},
			args: struct{ p2 *Point }{p2: NewPoint(3, 1)},
			want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1 := &Point{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := p1.Distance(tt.args.p2); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
