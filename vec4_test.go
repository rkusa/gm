package ml

import (
	"reflect"
	"testing"
)

func TestVec4AddGo(t *testing.T) {
	testVec4Add(t, addVec4)
}

func TestVec4AddSIMD(t *testing.T) {
	testVec4Add(t, addVec4SIMD)
}

func TestVec4Add(t *testing.T) {
	testVec4Add(t, func(lhs, rhs *Vec4) {
		lhs.Add(rhs)
	})
}

func testVec4Add(t *testing.T, add func(lhs, rhs *Vec4)) {
	lhs := &Vec4{1, 2, 3, 4}
	rhs := &Vec4{5, 6, 7, 8}

	add(lhs, rhs)
	if !reflect.DeepEqual(lhs, &Vec4{6, 8, 10, 12}) {
		t.Fatalf("Add wrong result, got: %v", lhs)
	}

	// test add itself
	lhs = &Vec4{1, 2, 3, 4}

	add(lhs, lhs)
	if !reflect.DeepEqual(lhs, &Vec4{2, 4, 6, 8}) {
		t.Fatalf("Add itself wrong result, got: %v", lhs)
	}
}

func BenchmarkVec4AddGo(b *testing.B) {
	benchmarkVec4Add(b, addVec4)
}

func BenchmarkVec4AddSIMD(b *testing.B) {
	benchmarkVec4Add(b, addVec4SIMD)
}

func benchmarkVec4Add(b *testing.B, add func(lhs, rhs *Vec4)) {
	lhs := &Vec4{1, 2, 3, 4}
	rhs := &Vec4{5, 6, 7, 8}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		add(lhs, rhs)
	}
}

func TestVec4LenGo(t *testing.T) {
	testVec4Len(t, lenVec4)
}

func TestVec4LenSIMD(t *testing.T) {
	testVec4Len(t, lenVec4SIMD)
}

func TestVec4Len(t *testing.T) {
	testVec4Len(t, (*Vec4).Len)
}

func testVec4Len(t *testing.T, len func(lhs *Vec4) float32) {
	lhs := &Vec4{1, 2, 3, 4}

	l := len(lhs)
	if l != 5.477226 {
		t.Fatalf("Len wrong result, got: %v, %v", l, lhs)
	}
}

func BenchmarkVec4LenGo(b *testing.B) {
	benchmarkVec4Len(b, lenVec4)
}

func BenchmarkVec4LenSIMD(b *testing.B) {
	benchmarkVec4Len(b, lenVec4SIMD)
}

func benchmarkVec4Len(b *testing.B, len func(lhs *Vec4) float32) {
	lhs := &Vec4{1, 2, 3, 4}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		len(lhs)
	}
}

func TestVec4MulGo(t *testing.T) {
	testVec4Mul(t, mulVec4)
}

func TestVec4MulSIMD(t *testing.T) {
	testVec4Mul(t, mulVec4SIMD)
}

func TestVec4(t *testing.T) {
	testVec4Mul(t, func(lhs *Vec4, rhs float32) {
		lhs.Mul(rhs)
	})
}

func testVec4Mul(t *testing.T, mul func(lhs *Vec4, rhs float32)) {
	lhs := &Vec4{1, 2, 3, 4}
	var rhs float32 = 2.5
	lhs.Mul(rhs)
	if !reflect.DeepEqual(lhs, &Vec4{2.5, 5, 7.5, 10}) {
		t.Fatalf("Mul wrong result, got: %v", lhs)
	}
}

func BenchmarkVec4MulGo(b *testing.B) {
	benchmarkVec4Mul(b, mulVec4)
}

func BenchmarkVec4MulSIMD(b *testing.B) {
	benchmarkVec4Mul(b, mulVec4SIMD)
}

func benchmarkVec4Mul(b *testing.B, mul func(lhs *Vec4, rhs float32)) {
	lhs := &Vec4{1, 2, 3, 4}
	var rhs float32 = 2.5
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Mul(rhs)
	}
}
