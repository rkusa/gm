package vec4

import "github.com/rkusa/gm/mat4"

type Vec4 [4]float32

func New(x, y, z, w float32) *Vec4 {
	return &Vec4{x, y, z, w}
}

// Clone the vector.
func (lhs *Vec4) Clone() *Vec4 {
	return &Vec4{lhs[0], lhs[1], lhs[2], lhs[3]}
}

// Transform the vector using a 4x4 matrix. Returns itself for function
// chaining.
func (lhs *Vec4) Transform(rhs *mat4.Mat4) *Vec4 {
	x, y, z, w := lhs[0], lhs[1], lhs[2], lhs[3]
	lhs[0] = rhs[0]*x + rhs[4]*y + rhs[8]*z + rhs[12]*w
	lhs[1] = rhs[1]*x + rhs[5]*y + rhs[9]*z + rhs[13]*w
	lhs[2] = rhs[2]*x + rhs[6]*y + rhs[10]*z + rhs[14]*w
	lhs[3] = rhs[3]*x + rhs[7]*y + rhs[11]*z + rhs[15]*w
	return lhs
}
