package convexhull

import "math"

// Point ...
type Point struct {
	X float64
	Y float64
}

// LessThan ...
func (p1 Point) LessThan(p2 Point) bool {
	if p1.X != p2.X {
		return p1.X < p2.X
	}
	return p1.Y <= p2.Y
}

// Angle ...
func (p1 Point) Angle(p2 Point) float64 {
	dY := p2.Y - p1.Y
	dX := p2.X - p1.X
	return math.Atan2(dY, dX)
}
