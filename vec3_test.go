package ml

import (
	"reflect"
	"testing"
)

func TestVec3Cross(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	rhs := &Vec3{4, 5, 6}

	lhs.Cross(rhs)
	if !reflect.DeepEqual(lhs, &Vec3{-3, 6, -3}) {
		t.Fatalf("Cross wrong result, got: %v", lhs)
	}
}

func TestVec3LenGo(t *testing.T) {
	testVec3Len(t, lenVec3)
}

func TestVec3LenSIMD(t *testing.T) {
	testVec3Len(t, lenVec3SIMD)
}

func TestVec3Len(t *testing.T) {
	testVec3Len(t, (*Vec3).Len)
}

func testVec3Len(t *testing.T, len func(lhs *Vec3) float32) {
	lhs := &Vec3{1, 2, 3}
	length := len(lhs)

	if length != 3.7416575 {
		t.Fatalf("Len wrong result, got: %v", length)
	}
}

func BenchmarkVec3LenGo(b *testing.B) {
	benchmarkVec3Len(b, lenVec3)
}

func BenchmarkVec3LenSIMD(b *testing.B) {
	benchmarkVec3Len(b, lenVec3SIMD)
}

func benchmarkVec3Len(b *testing.B, len func(lhs *Vec3) float32) {
	lhs := &Vec3{1, 2, 3}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		len(lhs)
	}
}

func TestVec3Normalize(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	lhs.Normalize()

	if !reflect.DeepEqual(lhs, &Vec3{0.26726124, 0.5345225, 0.8017837}) {
		t.Fatalf("Normalize wrong result, got: %v", lhs)
	}
}

func TestVec3SubGo(t *testing.T) {
	testVec3Sub(t, subVec3)
}

func TestVec3SubSIMD(t *testing.T) {
	testVec3Sub(t, subVec3SIMD)
}

func TestVec3Sub(t *testing.T) {
	testVec3Sub(t, func(lhs, rhs *Vec3) {
		lhs.Sub(rhs)
	})
}

func testVec3Sub(t *testing.T, sub func(lhs, rhs *Vec3)) {
	lhs := &Vec3{1, 2, 3}
	rhs := &Vec3{4, 5, 6}

	sub(lhs, rhs)
	if !reflect.DeepEqual(lhs, &Vec3{-3, -3, -3}) {
		t.Fatalf("Sub wrong result, got: %v", lhs)
	}

	// test sub itself
	lhs = &Vec3{1, 2, 3}

	sub(lhs, lhs)
	if !reflect.DeepEqual(lhs, &Vec3{0, 0, 0}) {
		t.Fatalf("Sub itself wrong result, got: %v", lhs)
	}
}

func BenchmarkVec3SubGo(b *testing.B) {
	benchmarkVec3Sub(b, subVec3)
}

func BenchmarkVec3SubSIMD(b *testing.B) {
	benchmarkVec3Sub(b, subVec3SIMD)
}

func benchmarkVec3Sub(b *testing.B, sub func(lhs, rhs *Vec3)) {
	lhs := &Vec3{1, 2, 3}
	rhs := &Vec3{4, 5, 6}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		sub(lhs, rhs)
	}
}
