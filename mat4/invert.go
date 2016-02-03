package mat4

// Invert the matrix. Returns nil if inversion is not possible due to a zero
// row; otherwise, return itself for function chaining.
func (lhs *Mat4) Invert() *Mat4 {
	if !invertSIMD(lhs) {
		return nil
	}

	return lhs
}

func invert(out *Mat4) bool {
	lhs := Mat4{
		out[0], out[1], out[2], out[3],
		out[4], out[5], out[6], out[7],
		out[8], out[9], out[10], out[11],
		out[12], out[13], out[14], out[15],
	}

	s0 := lhs[0]*lhs[5] - lhs[4]*lhs[1]
	s1 := lhs[0]*lhs[6] - lhs[4]*lhs[2]
	s2 := lhs[0]*lhs[7] - lhs[4]*lhs[3]
	s3 := lhs[1]*lhs[6] - lhs[5]*lhs[2]
	s4 := lhs[1]*lhs[7] - lhs[5]*lhs[3]
	s5 := lhs[2]*lhs[7] - lhs[6]*lhs[3]

	c5 := lhs[10]*lhs[15] - lhs[14]*lhs[11]
	c4 := lhs[9]*lhs[15] - lhs[13]*lhs[11]
	c3 := lhs[9]*lhs[14] - lhs[13]*lhs[10]
	c2 := lhs[8]*lhs[15] - lhs[12]*lhs[11]
	c1 := lhs[8]*lhs[14] - lhs[12]*lhs[10]
	c0 := lhs[8]*lhs[13] - lhs[12]*lhs[9]

	det := s0*c5 - s1*c4 + s2*c3 + s3*c2 - s4*c1 + s5*c0
	if det == 0 {
		return false
	}

	out[0] = (lhs[5]*c5 - lhs[6]*c4 + lhs[7]*c3) / det
	out[1] = (-lhs[1]*c5 + lhs[2]*c4 - lhs[3]*c3) / det
	out[2] = (lhs[13]*s5 - lhs[14]*s4 + lhs[15]*s3) / det
	out[3] = (-lhs[9]*s5 + lhs[10]*s4 - lhs[11]*s3) / det

	out[4] = (-lhs[4]*c5 + lhs[6]*c2 - lhs[7]*c1) / det
	out[5] = (lhs[0]*c5 - lhs[2]*c2 + lhs[3]*c1) / det
	out[6] = (-lhs[12]*s5 + lhs[14]*s2 - lhs[15]*s1) / det
	out[7] = (lhs[8]*s5 - lhs[10]*s2 + lhs[11]*s1) / det

	out[8] = (lhs[4]*c4 - lhs[5]*c2 + lhs[7]*c0) / det
	out[9] = (-lhs[0]*c4 + lhs[1]*c2 - lhs[3]*c0) / det
	out[10] = (lhs[12]*s4 - lhs[13]*s2 + lhs[15]*s0) / det
	out[11] = (-lhs[8]*s4 + lhs[9]*s2 - lhs[11]*s0) / det

	out[12] = (-lhs[4]*c3 + lhs[5]*c1 - lhs[6]*c0) / det
	out[13] = (lhs[0]*c3 - lhs[1]*c1 + lhs[2]*c0) / det
	out[14] = (-lhs[12]*s3 + lhs[13]*s1 - lhs[14]*s0) / det
	out[15] = (lhs[8]*s3 - lhs[9]*s1 + lhs[10]*s0) / det

	return true
}
