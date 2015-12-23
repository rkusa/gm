package simd

type Vec4 [4]float32

func (lhs *Vec4) Add(rhs *Vec4) {
	addVec4(lhs, rhs)
}

func addVec4(lhs, rhs *Vec4) Vec4

func (lhs *Vec4) Mul(rhs float32) {
	mulVec4(lhs, rhs)
}

func mulVec4(lhs *Vec4, rhs float32)
