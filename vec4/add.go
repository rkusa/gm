package vec4

// Add two 4-dimensional vectors. Returns itself for function chaining.
func (lhs *Vec4) Add(rhs *Vec4) *Vec4 {
	addSIMD(lhs, rhs)
	return lhs
}

func addSIMD(lhs, rhs *Vec4)

func add(lhs, rhs *Vec4) {
	lhs[0] += rhs[0]
	lhs[1] += rhs[1]
	lhs[2] += rhs[2]
	lhs[3] += rhs[3]
}
