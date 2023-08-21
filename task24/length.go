package task24

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p1 *Point) Distance(p2 *Point) float64 {
	deltaX := p1.x - p2.x
	deltaXPow2 := math.Pow(deltaX, 2)

	deltaY := p1.y - p2.y
	deltaYPow2 := math.Pow(deltaY, 2)

	return math.Sqrt(deltaXPow2 + deltaYPow2)
}
