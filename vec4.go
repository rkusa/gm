package gm

import "github.com/rkusa/gm/math32"

type Vec4 [4]float32

func addVec4SIMD(lhs, rhs *Vec4)

func addVec4(lhs, rhs *Vec4) {
	lhs[0] += rhs[0]
	lhs[1] += rhs[1]
	lhs[2] += rhs[2]
	lhs[3] += rhs[3]
}

// Add two 4-dimensional vectors. Returns itself for function chaining.
func (lhs *Vec4) Add(rhs *Vec4) *Vec4 {
	addVec4SIMD(lhs, rhs)
	return lhs
}

// Clone the vector.
func (lhs *Vec4) Clone() *Vec4 {
	return &Vec4{lhs[0], lhs[1], lhs[2], lhs[3]}
}

func divVec4SIMD(lhs *Vec4, rhs float32)

func divVec4(lhs *Vec4, rhs float32) {
	lhs[0] /= rhs
	lhs[1] /= rhs
	lhs[2] /= rhs
	lhs[3] /= rhs
}

// Div divides the the calling vector by the provided one. The result is
// saved back into the calling vector. Returns itself for function chaining.
func (lhs *Vec4) Div(rhs float32) *Vec4 {
	divVec4(lhs, rhs)
	return lhs
}

func lenVec4SIMD(lhs *Vec4) float32

func lenVec4(lhs *Vec4) float32 {
	return math32.Sqrt(
		lhs[0]*lhs[0] + lhs[1]*lhs[1] + lhs[2]*lhs[2] + lhs[3]*lhs[3],
	)
}

// Len returns the vector length.
func (lhs *Vec4) Len() float32 {
	return lenVec4SIMD(lhs)
}

func mulVec4SIMD(lhs *Vec4, rhs float32)

func mulVec4(lhs *Vec4, rhs float32) {
	lhs[0] *= rhs
	lhs[1] *= rhs
	lhs[2] *= rhs
	lhs[3] *= rhs
}

// Multiply a 4-dimensional vector with a scalar. Returns itself for function
// chaining.
func (lhs *Vec4) Mul(rhs float32) *Vec4 {
	mulVec4(lhs, rhs)
	return lhs
}

func subVec4SIMD(lhs, rhs *Vec4)

func subVec4(lhs, rhs *Vec4) {
	lhs[0] -= rhs[0]
	lhs[1] -= rhs[1]
	lhs[2] -= rhs[2]
	lhs[3] -= rhs[3]
}

// Sub substracts two vectors. Returns itself for function chaining.
func (lhs *Vec4) Sub(rhs *Vec4) *Vec4 {
	subVec4(lhs, rhs)
	return lhs
}

// Transform the vector using a 4x4 matrix. Returns itself for function
// chaining.
func (lhs *Vec4) Transform(rhs *Mat4) *Vec4 {
	x, y, z, w := lhs[0], lhs[1], lhs[2], lhs[3]
	lhs[0] = rhs[0]*x + rhs[4]*y + rhs[8]*z + rhs[12]*w
	lhs[1] = rhs[1]*x + rhs[5]*y + rhs[9]*z + rhs[13]*w
	lhs[2] = rhs[2]*x + rhs[6]*y + rhs[10]*z + rhs[14]*w
	lhs[3] = rhs[3]*x + rhs[7]*y + rhs[11]*z + rhs[15]*w
	return lhs
}
