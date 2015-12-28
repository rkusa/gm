package gm

import (
	"reflect"
	"testing"

	"github.com/rkusa/gm/math32"
)

func TestMat4Identity(t *testing.T) {
	ident := &Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	m := Mat4Identity()
	if !reflect.DeepEqual(m, ident) {
		t.Fatalf("Mat4Identity wrong result, got: %v", m)
	}

	m = &Mat4{}
	m.Identity()
	if !reflect.DeepEqual(m, ident) {
		t.Fatalf("Identity wrong result, got: %v", m)
	}
}

func TestMat4InvertGo(t *testing.T) {
	testMat4Invert(t, invertMat4)
}

func TestMat4InvertSIMD(t *testing.T) {
	testMat4Invert(t, invertMat4SIMD)
}

func TestMat4Invert(t *testing.T) {
	testMat4Invert(t, func(lhs *Mat4) bool {
		return lhs.Invert() != nil
	})
}

func testMat4Invert(t *testing.T, invert func(lhs *Mat4) bool) {
	lhs := &Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 0,
	}
	clone := lhs.Clone()

	if invert(lhs) {
		t.Fatalf("Invert wrong result, does not break on 0 row")
	}

	if !reflect.DeepEqual(lhs, clone) {
		t.Fatalf("Invert should not mutate on 0 row")
	}

	lhs[15] = 1
	expectation := &Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		-1, -2, -3, 1,
	}

	if !invert(lhs) {
		t.Fatalf("Invert should return true if there is no 0 row")
	}

	if lhs == nil || !mat4Close(lhs, expectation) {
		t.Fatalf("Invert wrong result, got: %v", lhs)
	}
}

func BenchmarkMat4InvertGo(b *testing.B) {
	benchmarkMat4Invert(b, invertMat4)
}

func BenchmarkMat4InvertSIMD(b *testing.B) {
	benchmarkMat4Invert(b, invertMat4SIMD)
}

func benchmarkMat4Invert(b *testing.B, invert func(lhs *Mat4) bool) {
	lhs := Mat4Identity()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		invert(lhs)
	}
}

func TestMat4LookAt(t *testing.T) {
	eye, center, up := Vec3{3, 3, 3}, Vec3{0, 0, 0}, Vec3{0, 1, 0}
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

func BenchmarkMat4LookAt(b *testing.B) {
	eye, center, up := Vec3{3, 3, 3}, Vec3{0, 0, 0}, Vec3{0, 1, 0}
	m := &Mat4{}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.LookAt(&eye, center, up)
	}
}

func TestMat4MulGo(t *testing.T) {
	testMat4Mul(t, mulMat4)
}

func TestMat4MulSIMD(t *testing.T) {
	testMat4Mul(t, mulMat4SIMD)
}

func TestMat4Mul(t *testing.T) {
	testMat4Mul(t, func(lhs, rhs *Mat4) {
		lhs.Mul(rhs)
	})
}

func testMat4Mul(t *testing.T, mul func(lhs, rhs *Mat4)) {
	lhs := &Mat4{1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15, 4, 8, 12, 16}
	rhs := &Mat4{17, 21, 25, 29, 18, 22, 26, 30, 19, 23, 27, 31, 20, 24, 28, 32}

	mul(lhs, rhs)
	expectation := &Mat4{250, 618, 986, 1354, 260, 644, 1028, 1412, 270, 670, 1070, 1470, 280, 696, 1112, 1528}
	if !reflect.DeepEqual(lhs, expectation) {
		t.Fatalf("Add wrong result, got: %v", lhs)
	}

	// test mul itself
	lhs = &Mat4{1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15, 4, 8, 12, 16}

	mul(lhs, lhs)
	expectation = &Mat4{90, 202, 314, 426, 100, 228, 356, 484, 110, 254, 398, 542, 120, 280, 440, 600}
	if !reflect.DeepEqual(lhs, expectation) {
		t.Fatalf("Add wrong result, got: %v", lhs)
	}
}

func BenchmarkMat4MulGo(b *testing.B) {
	benchmarkMat4Mul(b, mulMat4)
}

func BenchmarkMat4MulSIMD(b *testing.B) {
	benchmarkMat4Mul(b, mulMat4SIMD)
}

func benchmarkMat4Mul(b *testing.B, mul func(lhs, rhs *Mat4)) {
	lhs := &Mat4{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	rhs := &Mat4{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		mul(lhs, rhs)
	}
}

func TestMat4Perspective(t *testing.T) {
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

func BenchmarkMat4Perspective(b *testing.B) {
	m := Mat4{}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.Perspective(math32.Pi/4, 1920.0/1080, .1, 100)
	}
}

func TestMat4Rotate(t *testing.T) {
	lhs := &Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 1,
	}
	rad := math32.Pi / 2
	lhs.Rotate(rad, &Vec3{1, 0, 0})

	expectation := &Mat4{
		0.99999994, 0, 0, 0,
		0, math32.Cos(rad), math32.Sin(rad), 0,
		0, -math32.Sin(rad), math32.Cos(rad), 0,
		1, 2, 3, 1,
	}
	if !reflect.DeepEqual(lhs, expectation) {
		t.Fatalf("Rotate wrong result, got: %v %v", lhs, expectation)
	}
}

func BenchmarkMat4Rotate(b *testing.B) {
	m := Mat4Identity()
	rad := math32.Pi / 2
	axis := &Vec3{1, 0, 0}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.Rotate(rad, axis)
	}
}

func TestMat4Translate(t *testing.T) {
	lhs := &Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 1,
	}
	lhs.Translate(&Vec3{4, 5, 6})

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

func BenchmarkMat4Translate(b *testing.B) {
	m := &Mat4{}
	v := &Vec3{}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.Translate(v)
	}
}

func TestMat4Transpose(t *testing.T) {
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

func BenchmarkMat4Transpose(b *testing.B) {
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

func close(a, b float32) bool {
	return tolerance(a, b, 4e-4)
}

func mat4Close(lhs, rhs *Mat4) bool {
	for i := 0; i < 16; i++ {
		if !close(lhs[i], rhs[i]) {
			return false
		}
	}

	return true
}
