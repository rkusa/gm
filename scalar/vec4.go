package scalar

type Vec4 [4]float32

func (lhs *Vec4) Add(rhs *Vec4) {
	lhs[0] += rhs[0]
	lhs[1] += rhs[1]
	lhs[2] += rhs[2]
	lhs[3] += rhs[3]
}

func (lhs *Vec4) Mul(rhs float32) {
	lhs[0] *= rhs
	lhs[1] *= rhs
	lhs[2] *= rhs
	lhs[3] *= rhs
}
