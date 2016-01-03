package mat4

import (
	"reflect"
	"testing"

	"github.com/rkusa/gm/math32"
	"github.com/rkusa/gm/vec3"
)

func TestNew(t *testing.T) {
	lhs := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	expect := &Mat4{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	if !reflect.DeepEqual(lhs, expect) {
		t.Fatalf("New wrong result, got %v", lhs)
	}
}

func TestClone(t *testing.T) {
	a := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	b := a.Clone()

	if a == b {
		t.Fatalf("Clone must create a new instance")
	}

	expect := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	if !reflect.DeepEqual(a, expect) {
		t.Fatalf("Clone must not change values")
	}

	if !reflect.DeepEqual(b, expect) {
		t.Fatalf("Clone must keep values")
	}
}

func TestCopy(t *testing.T) {
	lhs := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	rhs := New(17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32)

	lhs.Copy(rhs)

	if lhs == rhs {
		t.Fatalf("Copy must keep seperate instances")
	}

	if !reflect.DeepEqual(rhs, lhs) {
		t.Fatalf("Copy must keep values, got %v", rhs)
	}
}

func TestIdentity(t *testing.T) {
	ident := &Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	m := Identity()
	if !reflect.DeepEqual(m, ident) {
		t.Fatalf("Identity wrong result, got: %v", m)
	}

	m = &Mat4{}
	m.Identity()
	if !reflect.DeepEqual(m, ident) {
		t.Fatalf("Identity wrong result, got: %v", m)
	}
}

func TestLookAt(t *testing.T) {
	eye, center, up := vec3.Vec3{3, 3, 3}, vec3.Vec3{0, 0, 0}, vec3.Vec3{0, 1, 0}
	m := &Mat4{}
	m.LookAt(&eye, center, up)

	expectation := &Mat4{
		0.70710677, -0.4082483, 0.5773503, 0,
		0, 0.8164966, 0.5773503, 0,
		-0.70710677, -0.4082483, 0.5773503, 0,
		0, 0, -5.1961527, 1,
	}
	if !reflect.DeepEqual(m, expectation) {
		t.Fatalf("Translate wrong result, got: %v", m)
	}
}

func BenchmarkLookAt(b *testing.B) {
	eye, center, up := vec3.Vec3{3, 3, 3}, vec3.Vec3{0, 0, 0}, vec3.Vec3{0, 1, 0}
	m := &Mat4{}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.LookAt(&eye, center, up)
	}
}

func TestPerspective(t *testing.T) {
	m := Mat4{}
	m.Perspective(math32.Pi/4, 1920.0/1080, .1, 100)

	expectation := Mat4{
		1.357995, 0, 0, 0,
		0, 2.4142134, 0, 0,
		0, 0, -1.002002, -1,
		0, 0, -0.2002002, 0,
	}
	if !reflect.DeepEqual(m, expectation) {
		t.Fatalf("Perspective wrong result, got: %v", m)
	}
}

func BenchmarkPerspective(b *testing.B) {
	m := Mat4{}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.Perspective(math32.Pi/4, 1920.0/1080, .1, 100)
	}
}

func TestRotation(t *testing.T) {
	mat := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	theta := math32.Pi / 2
	rot := vec3.New(theta, 0, 0)
	a := math32.Cos(theta)
	b := math32.Sin(theta)

	x := &Mat4{
		1, 0, 0, 0,
		0, a, b, 0,
		0, -b, a, 0,
		0, 0, 0, 1,
	}
	lhs := mat.Clone()
	lhs.Rotation(rot)

	if !reflect.DeepEqual(lhs, x) {
		t.Fatalf("Rotation X wrong result, got: %v %v", lhs, x)
	}

	y := &Mat4{
		a, 0, -b, 0,
		0, 1, 0, 0,
		b, 0, a, 0,
		0, 0, 0, 1,
	}
	lhs = mat.Clone()
	rot = vec3.New(0, theta, 0)
	lhs.Rotation(rot)

	if !reflect.DeepEqual(lhs, y) {
		t.Fatalf("Rotation Y wrong result,\ngot:\n%vexpected:\n%v", lhs, y)
	}

	z := &Mat4{
		a, b, 0, 0,
		-b, a, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	lhs = mat.Clone()
	rot = vec3.New(0, 0, theta)
	lhs.Rotation(rot)

	if !reflect.DeepEqual(lhs, z) {
		t.Fatalf("Rotation Z wrong result,\ngot:\n%vexpected:\n%v", lhs, z)
	}
}

func BenchmarkRotation(b *testing.B) {
	m := Identity()
	rot := vec3.New(math32.Pi/2, 0, 0)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.Rotation(rot)
	}
}

func TestTranslate(t *testing.T) {
	lhs := &Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 1,
	}
	lhs.Translate(&vec3.Vec3{4, 5, 6})

	expectation := &Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		5, 7, 9, 1,
	}
	if !reflect.DeepEqual(lhs, expectation) {
		t.Fatalf("Translate wrong result, got: %v", lhs)
	}
}

func BenchmarkTranslate(b *testing.B) {
	m := &Mat4{}
	v := &vec3.Vec3{}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.Translate(v)
	}
}

func TestTranspose(t *testing.T) {
	lhs := &Mat4{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	expectation := &Mat4{
		1, 5, 9, 13,
		2, 6, 10, 14,
		3, 7, 11, 15,
		4, 8, 12, 16,
	}
	lhs.Transpose()

	if !reflect.DeepEqual(lhs, expectation) {
		t.Fatalf("Transpose wrong result, got: %v", lhs)
	}
}

func BenchmarkTranspose(b *testing.B) {
	m := &Mat4{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.Transpose()
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

func close(lhs, rhs *Mat4) bool {
	for i := 0; i < 16; i++ {
		if !tolerance(lhs[i], rhs[i], 4e-4) {
			return false
		}
	}

	return true
}
