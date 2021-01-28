package geometry

import "math"

// Point defines a cartesian point in XY plane
type Point struct {
	X, Y float64
}

// Path defines a journe connecting the points with straight line
type Path []Point

// Distance function returns cartesian distance between two points
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance function returns cartesian distance between two points
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance returns the distance travelled along the path
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// ScaleBy scales a point
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
