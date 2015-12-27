package ml

import "github.com/rkusa/ml/math32"

type Mat4 [16]float32

func mulMat4SIMD(lhs, rhs *Mat4)

func mulMat4(out, rhs *Mat4) {
	// lhs := Mat4{out...}
	lhs := Mat4{
		out[0], out[1], out[2], out[3],
		out[4], out[5], out[6], out[7],
		out[8], out[9], out[10], out[11],
		out[12], out[13], out[14], out[15],
	}

	// if multiplicated with itself
	if out == rhs {
		rhs = &lhs
	}

	out[0] = lhs[0]*rhs[0] + lhs[4]*rhs[1] + lhs[8]*rhs[2] + lhs[12]*rhs[3]
	out[1] = lhs[1]*rhs[0] + lhs[5]*rhs[1] + lhs[9]*rhs[2] + lhs[13]*rhs[3]
	out[2] = lhs[2]*rhs[0] + lhs[6]*rhs[1] + lhs[10]*rhs[2] + lhs[14]*rhs[3]
	out[3] = lhs[3]*rhs[0] + lhs[7]*rhs[1] + lhs[11]*rhs[2] + lhs[15]*rhs[3]
	out[4] = lhs[0]*rhs[4] + lhs[4]*rhs[5] + lhs[8]*rhs[6] + lhs[12]*rhs[7]
	out[5] = lhs[1]*rhs[4] + lhs[5]*rhs[5] + lhs[9]*rhs[6] + lhs[13]*rhs[7]
	out[6] = lhs[2]*rhs[4] + lhs[6]*rhs[5] + lhs[10]*rhs[6] + lhs[14]*rhs[7]
	out[7] = lhs[3]*rhs[4] + lhs[7]*rhs[5] + lhs[11]*rhs[6] + lhs[15]*rhs[7]
	out[8] = lhs[0]*rhs[8] + lhs[4]*rhs[9] + lhs[8]*rhs[10] + lhs[12]*rhs[11]
	out[9] = lhs[1]*rhs[8] + lhs[5]*rhs[9] + lhs[9]*rhs[10] + lhs[13]*rhs[11]
	out[10] = lhs[2]*rhs[8] + lhs[6]*rhs[9] + lhs[10]*rhs[10] + lhs[14]*rhs[11]
	out[11] = lhs[3]*rhs[8] + lhs[7]*rhs[9] + lhs[11]*rhs[10] + lhs[15]*rhs[11]
	out[12] = lhs[0]*rhs[12] + lhs[4]*rhs[13] + lhs[8]*rhs[14] + lhs[12]*rhs[15]
	out[13] = lhs[1]*rhs[12] + lhs[5]*rhs[13] + lhs[9]*rhs[14] + lhs[13]*rhs[15]
	out[14] = lhs[2]*rhs[12] + lhs[6]*rhs[13] + lhs[10]*rhs[14] + lhs[14]*rhs[15]
	out[15] = lhs[3]*rhs[12] + lhs[7]*rhs[13] + lhs[11]*rhs[14] + lhs[15]*rhs[15]
}

// Multiplies two 4x4 matrices (using SIMD). Returns itself for function
// chaining.
func (lhs *Mat4) Mul(rhs *Mat4) *Mat4 {
	mulMat4SIMD(lhs, rhs)
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
