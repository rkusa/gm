package gm

type Mat3 [9]float32

// NormalMatrix calculates the normal Matrix (the inverse transpose). Returns
// itself for function chaining.
func (lhs *Mat3) NormalMatrix(rhs *Mat4) *Mat3 {
	rhs = rhs.Clone()
	rhs.Invert().Transpose()

	lhs[0], lhs[1], lhs[2] = rhs[0], rhs[1], rhs[2]
	lhs[3], lhs[4], lhs[5] = rhs[4], rhs[5], rhs[6]
	lhs[6], lhs[7], lhs[8] = rhs[8], rhs[9], rhs[10]

	return lhs
}
