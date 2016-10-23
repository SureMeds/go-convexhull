# go-convexhull

This package arranges a set of points in a convex hull

```go
import (
  "github.com/SureMeds/go-convexhull"
)

points := convexhull.Points{
  convexhull.Point{X: 10, Y: 10},
  convexhull.Point{X: 10, Y: 20},
  convexhull.Point{X: 20, Y: 10},
  convexhull.Point{X: 20, Y: 20}}

(&points).GrahamScan()
```
