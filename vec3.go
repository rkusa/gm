package ml

import "github.com/rkusa/ml/math32"

type Vec3 [3]float32

// Clone initializes a new Vec3 initialized with values from an existing one.
func (lhs *Vec3) Clone() *Vec3 {
	return &Vec3{lhs[0], lhs[1], lhs[2]}
}

// Cross calculates the vector cross product. Saves the result into the
// calling vector. Returns itself for function chaining.
func (lhs *Vec3) Cross(rhs *Vec3) *Vec3 {
	a, b, c := lhs[0], lhs[1], lhs[2]

	lhs[0] = b*rhs[2] - c*rhs[1]
	lhs[1] = c*rhs[0] - a*rhs[2]
	lhs[2] = a*rhs[1] - b*rhs[0]

	return lhs
}

// Div divides the the calling vector by the provided one. The result is
// saved back into the calling vector. Returns itself for function chaining.
func (lhs *Vec3) Div(rhs float32) *Vec3 {
	lhs[0] /= rhs
	lhs[1] /= rhs
	lhs[2] /= rhs
	return lhs
}

// Len returns the vector length.
func (lhs *Vec3) Len() float32 {
	return math32.Sqrt(lhs[0]*lhs[0] + lhs[1]*lhs[1] + lhs[2]*lhs[2])
}

// Norrmalize the vector. Returns itself for function chaining.
func (lhs *Vec3) Normalize() *Vec3 {
	lhs.Div(lhs.Len())
	return lhs
}

// Sub subtracts the provided vector from the calling one. The result is
// saved into the calling vector. Returns itself for function chaining.
func (lhs *Vec3) Sub(rhs *Vec3) *Vec3 {
	lhs[0] -= rhs[0]
	lhs[1] -= rhs[1]
	lhs[2] -= rhs[2]
	return lhs
}
