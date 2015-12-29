package vec4

// Sub substracts two vectors. Returns itself for function chaining.
func (lhs *Vec4) Sub(rhs *Vec4) *Vec4 {
	sub(lhs, rhs)
	return lhs
}

func subSIMD(lhs, rhs *Vec4)

func sub(lhs, rhs *Vec4) {
	lhs[0] -= rhs[0]
	lhs[1] -= rhs[1]
	lhs[2] -= rhs[2]
	lhs[3] -= rhs[3]
}
