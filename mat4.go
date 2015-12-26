package ml

import "github.com/rkusa/ml/math32"

type Mat4 [16]float32

// Multiplies two 4x4 matrices (using SIMD). Returns itself for function
// chaining.
func (lhs *Mat4) Mul(rhs *Mat4) *Mat4 {
	mat4MulSIMD(lhs, rhs)
	return lhs
}

// LookAt calculates a look-at matrix with the given eye position, focal point,
// and up axis. The result is saved into the calling matrix. Returns itself
// for function chaining.
func (lhs *Mat4) LookAt(eye *Vec3, center, up Vec3) *Mat4 {
	z := center.Sub(eye).Normalize()
	x := z.Clone().Cross(up.Normalize()).Normalize()
	y := x.Clone().Cross(z)

	lhs[0], lhs[1], lhs[2], lhs[3] = x[0], y[0], -z[0], 0
	lhs[4], lhs[5], lhs[6], lhs[7] = x[1], y[1], -z[1], 0
	lhs[8], lhs[9], lhs[10], lhs[11] = x[2], y[2], -z[2], 0
	lhs[12], lhs[13], lhs[14], lhs[15] = 0, 0, 0, 1

	lhs.Translate(&Vec3{-eye[0], -eye[1], -eye[2]})

	return lhs
}

// Generates a perspective projection matrix using the vertical field of view
// (fovy; in radians), the aspect radio (width/height) and the near and far
// frustum bounds. Returns itself for function chaining.
func (lhs *Mat4) Perspective(fovy, aspect, near, far float32) *Mat4 {
	f, nf := 1/math32.Tan(fovy/2), near-far
	lhs[0] = f / aspect
	lhs[1], lhs[2], lhs[3], lhs[4] = 0, 0, 0, 0
	lhs[5] = f
	lhs[6], lhs[7], lhs[8], lhs[9] = 0, 0, 0, 0
	lhs[10] = (far + near) / nf
	lhs[11], lhs[12], lhs[13] = -1, 0, 0
	lhs[14] = (2 * far * near) / nf
	lhs[15] = 0
	return lhs
}

// Translate the matrix by the given vector. Returns itself for function
// chaining.
func (lhs *Mat4) Translate(v *Vec3) *Mat4 {
	lhs[12] = lhs[0]*v[0] + lhs[4]*v[1] + lhs[8]*v[2] + lhs[12]
	lhs[13] = lhs[1]*v[0] + lhs[5]*v[1] + lhs[9]*v[2] + lhs[13]
	lhs[14] = lhs[2]*v[0] + lhs[6]*v[1] + lhs[10]*v[2] + lhs[14]
	lhs[15] = lhs[3]*v[0] + lhs[7]*v[1] + lhs[11]*v[2] + lhs[15]
	return lhs
}
