package mat3

import "github.com/rkusa/gm/mat4"

type Mat3 [9]float32

func New(a00, a01, a02, a10, a11, a12, a20, a21, a22 float32) *Mat3 {
	return &Mat3{
		a00, a01, a02,
		a10, a11, a12,
		a20, a21, a22,
	}
}

// Clone the matrix. Returns itself for function chaining.
func (lhs *Mat3) Clone() *Mat3 {
	return &Mat3{
		lhs[0], lhs[1], lhs[2],
		lhs[3], lhs[4], lhs[5],
		lhs[6], lhs[7], lhs[8],
	}
}

// Copy values from the given matrix. Returns itself for function chaining.
func (lhs *Mat3) Copy(rhs *Mat3) *Mat3 {
	lhs[0], lhs[1], lhs[2] = rhs[0], rhs[1], rhs[2]
	lhs[3], lhs[4], lhs[5] = rhs[3], rhs[4], rhs[5]
	lhs[6], lhs[7], lhs[8] = rhs[6], rhs[7], rhs[8]

	return lhs
}

// Create a new identity matrix.
func Identity() *Mat3 {
	return &Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

// Identity resets itself to the identity matrix. Returns itself for
// function chaining.
func (lhs *Mat3) Identity() *Mat3 {
	lhs[0], lhs[1], lhs[2] = 1, 0, 0
	lhs[3], lhs[4], lhs[5] = 0, 1, 0
	lhs[6], lhs[7], lhs[8] = 0, 0, 1
	return lhs
}

// NormalMatrix calculates the normal Matrix (the inverse transpose). Returns
// itself for function chaining.
func (lhs *Mat3) NormalMatrix(rhs *mat4.Mat4) *Mat3 {
	rhs = rhs.Clone()
	rhs.Invert().Transpose()

	lhs[0], lhs[1], lhs[2] = rhs[0], rhs[1], rhs[2]
	lhs[3], lhs[4], lhs[5] = rhs[4], rhs[5], rhs[6]
	lhs[6], lhs[7], lhs[8] = rhs[8], rhs[9], rhs[10]

	return lhs
}
