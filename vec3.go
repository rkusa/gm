package ml

import "github.com/rkusa/ml/math32"

type Vec3 [3]float32

func (lhs *Vec3) Cross() {

}

func (lhs *Vec3) Len() float32 {
	return math32.Sqrt(lhs[0]*lhs[0] + lhs[1]*lhs[1] + lhs[2]*lhs[2])
}

func (lhs *Vec3) Normalize() {
	l := lhs.Len()
	lhs[0] /= l
	lhs[1] /= l
	lhs[2] /= l
}

// Sub two 3-dimensional vectors
func (lhs *Vec3) Sub(rhs *Vec3) {
	vec3SubSIMD(lhs, rhs)
}
