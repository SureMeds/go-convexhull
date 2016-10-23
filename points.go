package convexhull

import "sort"

// Points ...
type Points []Point

//PointsByLocation ...
type PointsByLocation Points

//PointsByAngle ...
type PointsByAngle Points

// SortByLocation ...
func (p *Points) SortByLocation() {
	copyPoints := make(PointsByLocation, len(*p))
	for _, point := range *p {
		copyPoints = append(copyPoints, point)
	}
	sort.Sort(copyPoints)
	for i, point := range copyPoints {
		(*p)[i] = point
	}
}

// SortByAngle ...
func (p *Points) SortByAngle() {
	copyPoints := make(PointsByAngle, len(*p))
	for _, point := range *p {
		copyPoints = append(copyPoints, point)
	}
	sort.Sort(copyPoints)
	for i, point := range copyPoints {
		(*p)[i] = point
	}
}

// Len ...
func (p PointsByLocation) Len() int {
	return len(p)
}

// Swap ...
func (p PointsByLocation) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less ...
func (p PointsByLocation) Less(i, j int) bool {
	if p[i].Y < p[j].Y {
		return true
	} else if p[i].Y == p[j].Y {
		return p[i].X <= p[j].X
	}
	return false
}

// Len ...
func (p PointsByAngle) Len() int {
	return len(p)
}

// Swap ...
func (p PointsByAngle) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less ...
func (p PointsByAngle) Less(i, j int) bool {
	thetai := p[0].Angle(p[i])
	thetaj := p[0].Angle(p[j])
	if thetai < thetaj {
		return true
	}
	return false
}

// RemoveDuplicates ...
func (p *Points) RemoveDuplicates() {
	copyPoints := make(Points, len(*p))
	for _, point := range *p {
		unique := true
		for _, existingPoint := range copyPoints {
			unique = unique && !(point.X == existingPoint.X && point.Y == existingPoint.Y)
		}
		if unique {
			copyPoints = append(copyPoints, point)
		}
	}
	*p = copyPoints
}

// FindLowest ...
func (p *Points) FindLowest() {
	copyPoints := make(Points, len(*p))
	var index int
	for i, point := range *p {
		if point.LessThan((*p)[index]) {
			index = i
		}
	}
	copyPoints = append(copyPoints, (*p)[index])
	for i, point := range *p {
		if i != index {
			copyPoints = append(copyPoints, point)
		}
	}
	*p = copyPoints
}
