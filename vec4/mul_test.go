package vec4

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
	testMul(t, func(lhs *Vec4, rhs float32) {
		lhs.Mul(rhs)
	})
}

func testMul(t *testing.T, mul func(lhs *Vec4, rhs float32)) {
	lhs := &Vec4{1, 2, 3, 4}
	var rhs float32 = 2.5
	mul(lhs, rhs)

	if !reflect.DeepEqual(lhs, &Vec4{2.5, 5, 7.5, 10}) {
		t.Fatalf("Mul wrong result, got: %v", lhs)
	}
}

func BenchmarkMulGo(b *testing.B) {
	benchmarkMul(b, mul)
}

func BenchmarkMulSIMD(b *testing.B) {
	benchmarkMul(b, mulSIMD)
}

func benchmarkMul(b *testing.B, mul func(lhs *Vec4, rhs float32)) {
	lhs := &Vec4{1, 2, 3, 4}
	var rhs float32 = 2.5
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		mul(lhs, rhs)
	}
}
