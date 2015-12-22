package simd

type Vec4 [4]float32

func (lhs *Vec4) Add(rhs *Vec4) {
	add(lhs, rhs)
}

func add(lhs, rhs *Vec4) Vec4

func (lhs *Vec4) Mul(rhs float32) {
	mul(lhs, rhs)
}

func mul(lhs *Vec4, rhs float32)
