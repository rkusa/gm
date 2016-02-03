// +build 386

package mat4

// Invert the matrix. Returns nil if inversion is not possible due to a zero
// row; otherwise, return itself for function chaining.
func (lhs *Mat4) Invert() *Mat4 {
	tmp := New()

	if !invertSIMD(lhs, tmp) {
		return nil
	}

	return lhs
}

func invertSIMD(out, tmp *Mat4) bool
