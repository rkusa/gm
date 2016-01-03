package mat4

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/rkusa/gm/math32"
	"github.com/rkusa/gm/vec3"
)

type Mat4 [16]float32

func New(a00, a01, a02, a03, a10, a11, a12, a13, a20, a21, a22, a23, a30, a31, a32, a33 float32) *Mat4 {
	return &Mat4{
		a00, a01, a02, a03,
		a10, a11, a12, a13,
		a20, a21, a22, a23,
		a30, a31, a32, a33,
	}
}

// Clone the matrix. Returns itself for function chaining.
func (lhs *Mat4) Clone() *Mat4 {
	return &Mat4{
		lhs[0], lhs[1], lhs[2], lhs[3],
		lhs[4], lhs[5], lhs[6], lhs[7],
		lhs[8], lhs[9], lhs[10], lhs[11],
		lhs[12], lhs[13], lhs[14], lhs[15],
	}
}

// Copy values from the given matrix. Returns itself for function chaining.
func (lhs *Mat4) Copy(rhs *Mat4) *Mat4 {
	lhs[0], lhs[1], lhs[2], lhs[3] = rhs[0], rhs[1], rhs[2], rhs[3]
	lhs[4], lhs[5], lhs[6], lhs[7] = rhs[4], rhs[5], rhs[6], rhs[7]
	lhs[8], lhs[9], lhs[10], lhs[11] = rhs[8], rhs[9], rhs[10], rhs[11]
	lhs[12], lhs[13], lhs[14], lhs[15] = rhs[12], rhs[13], rhs[14], rhs[15]

	return lhs
}

// Create a new identity matrix.
func Identity() *Mat4 {
	return &Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Identity resets itself to the identity matrix. Returns itself for
// function chaining.
func (lhs *Mat4) Identity() *Mat4 {
	lhs[0], lhs[1], lhs[2], lhs[3] = 1, 0, 0, 0
	lhs[4], lhs[5], lhs[6], lhs[7] = 0, 1, 0, 0
	lhs[8], lhs[9], lhs[10], lhs[11] = 0, 0, 1, 0
	lhs[12], lhs[13], lhs[14], lhs[15] = 0, 0, 0, 1
	return lhs
}

// LookAt calculates a look-at matrix with the given eye position, focal point,
// and up axis. The result is saved into the calling matrix. Returns itself
// for function chaining.
func (lhs *Mat4) LookAt(eye *vec3.Vec3, center, up vec3.Vec3) *Mat4 {
	z := center.Sub(eye).Normalize()
	x := z.Clone().Cross(up.Normalize()).Normalize()
	y := x.Clone().Cross(z)

	lhs[0], lhs[1], lhs[2], lhs[3] = x[0], y[0], -z[0], 0
	lhs[4], lhs[5], lhs[6], lhs[7] = x[1], y[1], -z[1], 0
	lhs[8], lhs[9], lhs[10], lhs[11] = x[2], y[2], -z[2], 0
	lhs[12], lhs[13], lhs[14], lhs[15] = 0, 0, 0, 1

	lhs.Translate(&vec3.Vec3{-eye[0], -eye[1], -eye[2]})

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

// Rotate the matrix by the given angles (x, y, z; in radians). Returns itself.
func (lhs *Mat4) Rotate(rot *vec3.Vec3) *Mat4 {
	var a, b, c, d, e, f float32

	if rot[0] != 0 {
		a = math32.Cos(rot[0])
		b = math32.Sin(rot[0])
	} else {
		a, b = 1, 0
	}

	if rot[1] != 0 {
		c = math32.Cos(rot[1])
		d = math32.Sin(rot[1])
	} else {
		c, d = 1, 0
	}

	if rot[2] != 0 {
		e = math32.Cos(rot[2])
		f = math32.Sin(rot[2])
	} else {
		e, f = 1, 0
	}

	ad, bd := a*d, b*d

	lhs[0], lhs[1], lhs[2] = c*e, bd*e+a*f, -ad*e+b*f
	lhs[4], lhs[5], lhs[6] = -c*f, -bd*f+a*e, ad*f+b*e
	lhs[8], lhs[9], lhs[10] = d, -b*c, a*c

	return lhs
}

// Translate the matrix by the given vector. Returns itself for function
// chaining.
func (lhs *Mat4) Translate(v *vec3.Vec3) *Mat4 {
	lhs[12] = lhs[0]*v[0] + lhs[4]*v[1] + lhs[8]*v[2] + lhs[12]
	lhs[13] = lhs[1]*v[0] + lhs[5]*v[1] + lhs[9]*v[2] + lhs[13]
	lhs[14] = lhs[2]*v[0] + lhs[6]*v[1] + lhs[10]*v[2] + lhs[14]
	lhs[15] = lhs[3]*v[0] + lhs[7]*v[1] + lhs[11]*v[2] + lhs[15]
	return lhs
}

// Transpose the matrix. Returns itself for function chaining.
func (lhs *Mat4) Transpose() *Mat4 {
	lhs[1], lhs[4] = lhs[4], lhs[1]
	lhs[2], lhs[8] = lhs[8], lhs[2]
	lhs[3], lhs[12] = lhs[12], lhs[3]
	lhs[6], lhs[9] = lhs[9], lhs[6]
	lhs[7], lhs[13] = lhs[13], lhs[7]
	lhs[11], lhs[14] = lhs[14], lhs[11]
	return lhs
}

// String pretty prints the matrix (for debugging purposes).
func (lhs Mat4) String() string {
	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 4, 4, 1, ' ', tabwriter.AlignRight)
	for i := 0; i < 4; i++ {
		fmt.Fprintf(w, "%v\t", lhs[i])
		fmt.Fprintf(w, "%v\t", lhs[i+4])
		fmt.Fprintf(w, "%v\t", lhs[i+8])
		fmt.Fprintf(w, "%v\t\n", lhs[i+12])
	}
	w.Flush()

	return buf.String()
}
