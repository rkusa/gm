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

func BenchmarkVec3Cross(b *testing.B) {
	lhs := &Vec3{1, 1, 1}
	rhs := &Vec3{1, 1, 1}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Cross(rhs)
	}
}

func TestVec3Div(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	var rhs float32 = 2

	lhs.Div(rhs)
	if !reflect.DeepEqual(lhs, &Vec3{.5, 1, 1.5}) {
		t.Fatalf("Div wrong result, got: %v", lhs)
	}
}

func BenchmarkVec3Div(b *testing.B) {
	lhs := &Vec3{1, 1, 1}
	var rhs float32 = 1
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Div(rhs)
	}
}

func TestVec3Len(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	length := lhs.Len()

	if length != 3.7416575 {
		t.Fatalf("Len wrong result, got: %v", length)
	}
}

func BenchmarkVec3Len(b *testing.B) {
	lhs := &Vec3{1, 2, 3}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Len()
	}
}

func TestVec3Normalize(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	lhs.Normalize()

	if !reflect.DeepEqual(lhs, &Vec3{0.26726124, 0.5345225, 0.8017837}) {
		t.Fatalf("Normalize wrong result, got: %v", lhs)
	}
}

func BenchmarkVec3Normalize(b *testing.B) {
	lhs := &Vec3{1, 2, 3}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Normalize()
	}
}

func TestVec3Sub(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	rhs := &Vec3{6, 5, 4}

	lhs.Sub(rhs)
	if !reflect.DeepEqual(lhs, &Vec3{-5, -3, -1}) {
		t.Fatalf("Sub wrong result, got: %v", lhs)
	}

	// test sub itself
	lhs = &Vec3{1, 2, 3}

	lhs.Sub(lhs)
	if !reflect.DeepEqual(lhs, &Vec3{0, 0, 0}) {
		t.Fatalf("Sub itself wrong result, got: %v", lhs)
	}
}

func BenchmarkVec3Sub(b *testing.B) {
	lhs := &Vec3{1, 2, 3}
	rhs := &Vec3{4, 5, 6}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Sub(rhs)
	}
}
