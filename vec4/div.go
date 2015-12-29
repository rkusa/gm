package vec4

// Div divides the the calling vector by the provided one. The result is
// saved back into the calling vector. Returns itself for function chaining.
func (lhs *Vec4) Div(rhs float32) *Vec4 {
	div(lhs, rhs)
	return lhs
}

func divSIMD(lhs *Vec4, rhs float32)

func div(lhs *Vec4, rhs float32) {
	lhs[0] /= rhs
	lhs[1] /= rhs
	lhs[2] /= rhs
	lhs[3] /= rhs
}
