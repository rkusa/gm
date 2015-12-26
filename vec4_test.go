package ml

import (
	"reflect"
	"testing"
)

func TestVec4AddScalar(t *testing.T) {
	testVec4Add(t, vec4AddScalar)
}

func TestVec4AddSIMD(t *testing.T) {
	testVec4Add(t, vec4AddSIMD)
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

func BenchmarkVec4AddScalar(b *testing.B) {
	benchmarkVec4Add(b, vec4AddScalar)
}

func BenchmarkVec4AddSIMD(b *testing.B) {
	benchmarkVec4Add(b, vec4AddSIMD)
}

func benchmarkVec4Add(b *testing.B, add func(lhs, rhs *Vec4)) {
	lhs := &Vec4{1, 2, 3, 4}
	rhs := &Vec4{5, 6, 7, 8}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		add(lhs, rhs)
	}
}

func TestVec4MulScalar(t *testing.T) {
	testVec4Mul(t, vec4MulScalar)
}

func TestVec4MulSIMD(t *testing.T) {
	testVec4Mul(t, vec4MulSIMD)
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

func BenchmarkVec4MulScalar(b *testing.B) {
	benchmarkVec4Mul(b, vec4MulScalar)
}

func BenchmarkVec4MulSIMD(b *testing.B) {
	benchmarkVec4Mul(b, vec4MulSIMD)
}

func benchmarkVec4Mul(b *testing.B, mul func(lhs *Vec4, rhs float32)) {
	lhs := &Vec4{1, 2, 3, 4}
	var rhs float32 = 2.5
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Mul(rhs)
	}
}
