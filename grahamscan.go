package convexhull

// CCW ...
// Three points are a counter-clockwise turn if ccw > 0, clockwise if
// ccw < 0, and collinear if ccw = 0 because ccw is a determinant that
// gives twice the signed  area of the triangle formed by p1, p2 and p3.
func CCW(p1, p2, p3 Point) bool {
	ccwID := (p2.X-p1.X)*(p3.Y-p1.Y) - (p2.Y-p1.Y)*(p3.X-p1.X)
	if ccwID < 0 {
		return false
	}
	return true
}

// GrahamScan ...
func (p *Points) GrahamScan() {
	// First sort by location
	p.SortByLocation()
	// Then sort by angle
	p.SortByAngle()
	// let N = number of points
	N := len(*p)
	// put last element as 0 element
	*p = append([]Point{(*p)[len(*p)-1]}, (*p)...)
	// M will denote the number of points on the convex hull.
	M := 1
	for i := 2; i < N; i++ {
		for CCW((*p)[M-1], (*p)[M], (*p)[i]) {
			if M > 1 {
				M = M - 1
			} else {
				i++
			}
		}
		M = M + 1
		// swap points[M] with points[i]
		temp := (*p)[M]
		(*p)[M] = (*p)[i]
		(*p)[i] = temp
	}
}
