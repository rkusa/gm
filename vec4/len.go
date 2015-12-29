package vec4

import "github.com/rkusa/gm/math32"

// Len returns the vector length.
func (lhs *Vec4) Len() float32 {
	return lenSIMD(lhs)
}

func lenSIMD(lhs *Vec4) float32

func len(lhs *Vec4) float32 {
	return math32.Sqrt(
		lhs[0]*lhs[0] + lhs[1]*lhs[1] + lhs[2]*lhs[2] + lhs[3]*lhs[3],
	)
}
