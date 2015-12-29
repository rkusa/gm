package mat4

import (
	"reflect"
	"testing"
)

func TestMulGo(t *testing.T) {
	testMul(t, mul)
}

func TestMulSIMD(t *testing.T) {
	testMul(t, mulSIMD)
}

func TestMul(t *testing.T) {
	testMul(t, func(lhs, rhs *Mat4) {
		lhs.Mul(rhs)
	})
}

func testMul(t *testing.T, mul func(lhs, rhs *Mat4)) {
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

func BenchmarkMulGo(b *testing.B) {
	benchmarkMul(b, mul)
}

func BenchmarkMulSIMD(b *testing.B) {
	benchmarkMul(b, mulSIMD)
}

func benchmarkMul(b *testing.B, mul func(lhs, rhs *Mat4)) {
	lhs := &Mat4{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	rhs := &Mat4{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		mul(lhs, rhs)
	}
}
