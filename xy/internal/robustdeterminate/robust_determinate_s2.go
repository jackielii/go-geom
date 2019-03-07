package robustdeterminate

import (
	"github.com/golang/geo/r3"
	"github.com/golang/geo/s2"
)

// SignOfDet2x2S2 computes the sign of the determinant of the 2x2 matrix
// with the given entries, in a robust way.
//
// return -1 if the determinant is negative,
//  return  1 if the determinant is positive,
// return  0 if the determinant is 0.
// It uses github.com/golang/geo/s2 under the hood
func SignOfDet2x2S2(x1, y1, x2, y2 float64) Sign {
	a := s2.Point{Vector: r3.Vector{0, 0, 0}}
	b := s2.Point{Vector: r3.Vector{x1, y1, 0}}
	c := s2.Point{Vector: r3.Vector{x2, y2, 0}}

	d := s2.RobustSign(a, b, c)
	return Sign(d)
}
