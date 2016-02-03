// +build !386

package mat4

// Invert the matrix. Returns nil if inversion is not possible due to a zero
// row; otherwise, return itself for function chaining.
func (lhs *Mat4) Invert() *Mat4 {
	if !invertSIMD(lhs) {
		return nil
	}

	return lhs
}

func invertSIMD(out *Mat4) bool
