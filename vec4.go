package ml

type Vec4 [4]float32

// Add two 4-dimensional vectors
func (lhs *Vec4) Add(rhs *Vec4) {
	vec4AddSIMD(lhs, rhs)
}

// Multiply a 4-dimensional vector with a scalar
func (lhs *Vec4) Mul(rhs float32) {
	vec4MulScalar(lhs, rhs)
}
