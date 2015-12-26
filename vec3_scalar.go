package ml

func vec3SubScalar(lhs, rhs *Vec3) {
	lhs[0] -= rhs[0]
	lhs[1] -= rhs[1]
	lhs[2] -= rhs[2]
}
