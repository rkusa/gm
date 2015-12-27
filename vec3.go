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

func lenVec3SIMD(lhs *Vec3) float32

func lenVec3(lhs *Vec3) float32 {
	return math32.Sqrt(lhs[0]*lhs[0] + lhs[1]*lhs[1] + lhs[2]*lhs[2])
}

// Len returns the vector length.
func (lhs *Vec3) Len() float32 {
	return lenVec3(lhs)
}

// Norrmalize the vector. Returns itself for function chaining.
func (lhs *Vec3) Normalize() *Vec3 {
	l := lhs.Len()
	lhs[0] /= l
	lhs[1] /= l
	lhs[2] /= l
	return lhs
}

func subVec3SIMD(lhs, rhs *Vec3)

func subVec3(lhs, rhs *Vec3) {
	lhs[0] -= rhs[0]
	lhs[1] -= rhs[1]
	lhs[2] -= rhs[2]
}

// Sub subtracts a the provided vector from the calling one. The result is
// saved into the calling vector. Returns itself for function chaining.
func (lhs *Vec3) Sub(rhs *Vec3) *Vec3 {
	subVec3SIMD(lhs, rhs)
	return lhs
}
