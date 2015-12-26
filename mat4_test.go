package ml

import (
	"github.com/rkusa/ml/math32"
	"reflect"
	"testing"
)

func TestMat4MulScalar(t *testing.T) {
	testMat4Mul(t, mat4MulScalar)
}

func TestMat4MulSIMD(t *testing.T) {
	testMat4Mul(t, mat4MulSIMD)
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

func BenchmarkMat4MulScalar(b *testing.B) {
	benchmarkMat4Mul(b, mat4MulScalar)
}

func BenchmarkMat4MulSIMD(b *testing.B) {
	benchmarkMat4Mul(b, mat4MulSIMD)
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
