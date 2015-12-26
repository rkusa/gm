package ml

import "github.com/rkusa/ml/math32"

type Vec3 [3]float32

func (lhs *Vec3) Len() float32 {
	return math32.Sqrt(lhs[0]*lhs[0] + lhs[1]*lhs[1] + lhs[2]*lhs[2])
}

// Sub two 3-dimensional vectors
func (lhs *Vec3) Sub(rhs *Vec3) {
	vec3SubSIMD(lhs, rhs)
}
