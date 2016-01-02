package mat3

import (
	"reflect"
	"testing"

	"github.com/rkusa/gm/mat4"
	"github.com/rkusa/gm/vec3"
)

func TestNew(t *testing.T) {
	lhs := New(1, 2, 3, 4, 5, 6, 7, 8, 9)
	expect := &Mat3{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(lhs, expect) {
		t.Fatalf("New wrong result, got %v", lhs)
	}
}

func TestClone(t *testing.T) {
	a := New(1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := a.Clone()

	if a == b {
		t.Fatalf("Clone must create a new instance")
	}

	expect := New(1, 2, 3, 4, 5, 6, 7, 8, 9)
	if !reflect.DeepEqual(a, expect) {
		t.Fatalf("Clone must not change values")
	}

	if !reflect.DeepEqual(b, expect) {
		t.Fatalf("Clone must keep values")
	}
}

func TestCopy(t *testing.T) {
	lhs := New(1, 2, 3, 4, 5, 6, 7, 8, 9)
	rhs := New(17, 18, 19, 20, 21, 22, 23, 24, 25)

	lhs.Copy(rhs)

	if lhs == rhs {
		t.Fatalf("Copy must keep seperate instances")
	}

	if !reflect.DeepEqual(rhs, lhs) {
		t.Fatalf("Copy must keep values, got %v", rhs)
	}
}

func TestIdentity(t *testing.T) {
	ident := &Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}

	m := Identity()
	if !reflect.DeepEqual(m, ident) {
		t.Fatalf("Identity wrong result, got: %v", m)
	}

	m = &Mat3{}
	m.Identity()
	if !reflect.DeepEqual(m, ident) {
		t.Fatalf("Identity wrong result, got: %v", m)
	}
}

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
