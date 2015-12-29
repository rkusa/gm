package mat3

import (
	"testing"

	"github.com/rkusa/gm/mat4"
	"github.com/rkusa/gm/vec3"
)

func TestNormalMatrix(t *testing.T) {
	p := mat4.Identity()
	p.Translate(&vec3.Vec3{2, 4, 6})
	// p.Rotate(math32.Pi/2, &vec3.Vec3{1, 0, 0})

	expectation := &Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}

	nm := &Mat3{}
	nm.NormalMatrix(p)

	if !close(nm, expectation) {
		t.Fatalf("NormalMatrix wrong result, got: %v", nm)
	}
}

func tolerance(a, b, e float32) bool {
	d := a - b
	if d < 0 {
		d = -d
	}

	// note: b is correct (expected) value, a is actual value.
	// make error tolerance a fraction of b, not a.
	if b != 0 {
		e = e * b
		if e < 0 {
			e = -e
		}
	}
	return d < e
}

func close(lhs, rhs *Mat3) bool {
	for i := 0; i < 9; i++ {
		if !tolerance(lhs[i], rhs[i], 4e-4) {
			return false
		}
	}

	return true
}
