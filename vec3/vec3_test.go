package vec3

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	lhs := New(1, 2, 3)

	if !reflect.DeepEqual(lhs, &Vec3{1, 2, 3}) {
		t.Fatalf("New wrong result, got %v", lhs)
	}
}

func TestClone(t *testing.T) {
	a := &Vec3{1, 2, 3}
	b := a.Clone()

	if a == b {
		t.Fatalf("Clone must create a new instance")
	}

	if !reflect.DeepEqual(a, &Vec3{1, 2, 3}) {
		t.Fatalf("Clone must not change values")
	}

	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Clone must keep values")
	}
}

func TestCross(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	rhs := &Vec3{4, 5, 6}

	lhs.Cross(rhs)
	if !reflect.DeepEqual(lhs, &Vec3{-3, 6, -3}) {
		t.Fatalf("Cross wrong result, got: %v", lhs)
	}
}

func BenchmarkCross(b *testing.B) {
	lhs := &Vec3{1, 1, 1}
	rhs := &Vec3{1, 1, 1}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Cross(rhs)
	}
}

func TestDiv(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	var rhs float32 = 2

	lhs.Div(rhs)
	if !reflect.DeepEqual(lhs, &Vec3{.5, 1, 1.5}) {
		t.Fatalf("Div wrong result, got: %v", lhs)
	}
}

func BenchmarkDiv(b *testing.B) {
	lhs := &Vec3{1, 1, 1}
	var rhs float32 = 1
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Div(rhs)
	}
}

func TestLen(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	length := lhs.Len()

	if length != 3.7416575 {
		t.Fatalf("Len wrong result, got: %v", length)
	}
}

func BenchmarkLen(b *testing.B) {
	lhs := &Vec3{1, 2, 3}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Len()
	}
}

func TestMul(t *testing.T) {
	lhs := &Vec3{2, 3, 4}
	var rhs float32 = 2

	lhs.Mul(rhs)
	if !reflect.DeepEqual(lhs, &Vec3{4, 6, 8}) {
		t.Fatalf("Mul wrong result, got: %v", lhs)
	}
}

func BenchmarkMul(b *testing.B) {
	lhs := &Vec3{1, 1, 1}
	var rhs float32 = 1
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Mul(rhs)
	}
}

func TestNormalize(t *testing.T) {
	lhs := &Vec3{1, 2, 3}
	lhs.Normalize()

	if !reflect.DeepEqual(lhs, &Vec3{0.26726124, 0.5345225, 0.8017837}) {
		t.Fatalf("Normalize wrong result, got: %v", lhs)
	}
}

func BenchmarkNormalize(b *testing.B) {
	lhs := &Vec3{1, 2, 3}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Normalize()
	}
}

func TestSub(t *testing.T) {
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

func BenchmarkSub(b *testing.B) {
	lhs := &Vec3{1, 2, 3}
	rhs := &Vec3{4, 5, 6}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Sub(rhs)
	}
}
