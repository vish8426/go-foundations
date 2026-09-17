package main

import "math"

// Point is a 2D position in metres.
type Point struct {
	X, Y float64
}

// Distance returns the Euclidean distance between p and q.
func Distance(p, q Point) float64 {
	dx := p.X - q.X
	dy := p.Y - q.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {}
