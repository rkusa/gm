package mat4

// Multiplies two 4x4 matrices (lhs*rhs; using SIMD). Saves result into out.
func Mul(out, lhs, rhs *Mat4) {
	mulSIMD(out, lhs, rhs)
}

// Multiplies two 4x4 matrices (using SIMD). Returns itself for function
// chaining.
func (lhs *Mat4) Mul(rhs *Mat4) *Mat4 {
	Mul(lhs, lhs, rhs)
	return lhs
}

func mulSIMD(out, lhs, rhs *Mat4)

func mul(out, lhs, rhs *Mat4) {
	if lhs == out && rhs == out {
		// if multiplicated with itself
		lhs = lhs.Clone()
		rhs = lhs
	} else if lhs == out {
		lhs = lhs.Clone()
	} else if rhs == out {
		rhs = rhs.Clone()
	}

	out[0] = lhs[0]*rhs[0] + lhs[4]*rhs[1] + lhs[8]*rhs[2] + lhs[12]*rhs[3]
	out[1] = lhs[1]*rhs[0] + lhs[5]*rhs[1] + lhs[9]*rhs[2] + lhs[13]*rhs[3]
	out[2] = lhs[2]*rhs[0] + lhs[6]*rhs[1] + lhs[10]*rhs[2] + lhs[14]*rhs[3]
	out[3] = lhs[3]*rhs[0] + lhs[7]*rhs[1] + lhs[11]*rhs[2] + lhs[15]*rhs[3]
	out[4] = lhs[0]*rhs[4] + lhs[4]*rhs[5] + lhs[8]*rhs[6] + lhs[12]*rhs[7]
	out[5] = lhs[1]*rhs[4] + lhs[5]*rhs[5] + lhs[9]*rhs[6] + lhs[13]*rhs[7]
	out[6] = lhs[2]*rhs[4] + lhs[6]*rhs[5] + lhs[10]*rhs[6] + lhs[14]*rhs[7]
	out[7] = lhs[3]*rhs[4] + lhs[7]*rhs[5] + lhs[11]*rhs[6] + lhs[15]*rhs[7]
	out[8] = lhs[0]*rhs[8] + lhs[4]*rhs[9] + lhs[8]*rhs[10] + lhs[12]*rhs[11]
	out[9] = lhs[1]*rhs[8] + lhs[5]*rhs[9] + lhs[9]*rhs[10] + lhs[13]*rhs[11]
	out[10] = lhs[2]*rhs[8] + lhs[6]*rhs[9] + lhs[10]*rhs[10] + lhs[14]*rhs[11]
	out[11] = lhs[3]*rhs[8] + lhs[7]*rhs[9] + lhs[11]*rhs[10] + lhs[15]*rhs[11]
	out[12] = lhs[0]*rhs[12] + lhs[4]*rhs[13] + lhs[8]*rhs[14] + lhs[12]*rhs[15]
	out[13] = lhs[1]*rhs[12] + lhs[5]*rhs[13] + lhs[9]*rhs[14] + lhs[13]*rhs[15]
	out[14] = lhs[2]*rhs[12] + lhs[6]*rhs[13] + lhs[10]*rhs[14] + lhs[14]*rhs[15]
	out[15] = lhs[3]*rhs[12] + lhs[7]*rhs[13] + lhs[11]*rhs[14] + lhs[15]*rhs[15]
}
