package ml

import "github.com/rkusa/ml/math32"

type Mat4 [16]float32

// Multiplies two 4x4 matrices (using SIMD)
func (lhs *Mat4) Mul(rhs *Mat4) {
	mat4MulSIMD(lhs, rhs)
}

// Generates a perspective projection matrix using the vertical field of view
// (fovy; in radians), the aspect radio (width/height) and the near and far
// frustum bounds.
func (out *Mat4) Perspective(fovy, aspect, near, far float32) {
	f, nf := 1/math32.Tan(fovy/2), near-far
	out[0] = f / aspect
	out[1], out[2], out[3], out[4] = 0, 0, 0, 0
	out[5] = f
	out[6], out[7], out[8], out[9] = 0, 0, 0, 0
	out[10] = (far + near) / nf
	out[11], out[12], out[13] = -1, 0, 0
	out[14] = (2 * far * near) / nf
	out[15] = 0
}

func (out *Mat4) LookAt(eye, center, up *Vec3) {
	// f := center.Sub(eye).Normalize()
	// s := f.Cross(up.Normalize()).Normalize()
	// u := s.Cross(f)

	// M := Mat4{
	// 	s[0], u[0], -f[0], 0,
	// 	s[1], u[1], -f[1], 0,
	// 	s[2], u[2], -f[2], 0,
	// 	0, 0, 0, 1,
	// }

	// return M.Mul4(Translate3D(float32(-eye[0]), float32(-eye[1]), float32(-eye[2])))
}
