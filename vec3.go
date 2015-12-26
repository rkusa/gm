package ml

import "github.com/rkusa/ml/math32"

type Vec3 [3]float32

// Cross calculates the vector cross product. Saves the result into the
// calling vector.
func (lhs *Vec3) Cross(rhs *Vec3) {
	a, b, c := lhs[0], lhs[1], lhs[2]

	lhs[0] = b*rhs[2] - c*rhs[1]
	lhs[1] = c*rhs[0] - a*rhs[2]
	lhs[2] = a*rhs[1] - b*rhs[0]
}

// Len returns the vector length.
func (lhs *Vec3) Len() float32 {
	return math32.Sqrt(lhs[0]*lhs[0] + lhs[1]*lhs[1] + lhs[2]*lhs[2])
}

// Norrmalize the vector.
func (lhs *Vec3) Normalize() {
	l := lhs.Len()
	lhs[0] /= l
	lhs[1] /= l
	lhs[2] /= l
}

// Sub subtracts a the provided vector from the calling one. The result is
// saved into the calling vector.
func (lhs *Vec3) Sub(rhs *Vec3) {
	vec3SubSIMD(lhs, rhs)
}
