package simd

type Mat4 [16]float32

func (lhs *Mat4) Mul(rhs *Mat4) {
	mulMat4(lhs, rhs)
}

func mulMat4(lhs, rhs *Mat4)
