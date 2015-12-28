package gm

import "testing"

func TestMat3NormalMatrix(t *testing.T) {
	p := Mat4Identity()
	p.Translate(&Vec3{2, 4, 6})
	// p.Rotate(math32.Pi/2, &Vec3{1, 0, 0})

	expectation := &Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}

	nm := &Mat3{}
	nm.NormalMatrix(p)

	if !mat3Close(nm, expectation) {
		t.Fatalf("NormalMatrix wrong result, got: %v", nm)
	}
}

func mat3Close(lhs, rhs *Mat3) bool {
	for i := 0; i < 9; i++ {
		if !close(lhs[i], rhs[i]) {
			return false
		}
	}

	return true
}
