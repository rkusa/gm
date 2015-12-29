package vec4

// Multiply a 4-dimensional vector with a scalar. Returns itself for function
// chaining.
func (lhs *Vec4) Mul(rhs float32) *Vec4 {
	mul(lhs, rhs)
	return lhs
}

func mulSIMD(lhs *Vec4, rhs float32)

func mul(lhs *Vec4, rhs float32) {
	lhs[0] *= rhs
	lhs[1] *= rhs
	lhs[2] *= rhs
	lhs[3] *= rhs
}
