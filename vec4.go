package ml

type Vec4 [4]float32

// Add two 4-dimensional vectors. Returns itself for function chaining.
func (lhs *Vec4) Add(rhs *Vec4) *Vec4 {
	vec4AddSIMD(lhs, rhs)
	return lhs
}

// Multiply a 4-dimensional vector with a scalar. Returns itself for function
// chaining.
func (lhs *Vec4) Mul(rhs float32) *Vec4 {
	vec4MulScalar(lhs, rhs)
	return lhs
}
