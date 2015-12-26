package ml

func vec4AddScalar(lhs, rhs *Vec4) {
	lhs[0] += rhs[0]
	lhs[1] += rhs[1]
	lhs[2] += rhs[2]
	lhs[3] += rhs[3]
}

func vec4MulScalar(lhs *Vec4, rhs float32) {
	lhs[0] *= rhs
	lhs[1] *= rhs
	lhs[2] *= rhs
	lhs[3] *= rhs
}
