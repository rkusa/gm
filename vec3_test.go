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

func TestVec3Len(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	length := lhs.Len()

	if length != 3.7416575 {
		t.Fatalf("Len wrong result, got: %v", length)
	}
}

func TestVec3Normalize(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	lhs.Normalize()

	if !reflect.DeepEqual(lhs, &Vec3{0.26726124, 0.5345225, 0.8017837}) {
		t.Fatalf("Normalize wrong result, got: %v", lhs)
	}
}

func TestVec3SubScalar(t *testing.T) {
	testVec3Sub(t, vec3SubScalar)
}

func TestVec3SubSIMD(t *testing.T) {
	testVec3Sub(t, vec3SubSIMD)
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

func BenchmarkVec3SubScalar(b *testing.B) {
	benchmarkVec3Sub(b, vec3SubScalar)
}

func BenchmarkVec3SubSIMD(b *testing.B) {
	benchmarkVec3Sub(b, vec3SubSIMD)
}

func benchmarkVec3Sub(b *testing.B, sub func(lhs, rhs *Vec3)) {
	lhs := &Vec3{1, 2, 3}
	rhs := &Vec3{4, 5, 6}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		sub(lhs, rhs)
	}
}
