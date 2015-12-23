package scalar

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	lhs := &Vec4{1, 2, 3, 4}
	rhs := &Vec4{5, 6, 7, 8}
	lhs.Add(rhs)
	if !reflect.DeepEqual(lhs, &Vec4{6, 8, 10, 12}) {
		t.Fatalf("Add wrong result, got: %v", lhs)
	}

	// test add itself
	lhs = &Vec4{1, 2, 3, 4}
	lhs.Add(lhs)
	if !reflect.DeepEqual(lhs, &Vec4{2, 4, 6, 8}) {
		t.Fatalf("Add itself wrong result, got: %v", lhs)
	}
}

func BenchmarkAdd(b *testing.B) {
	lhs := &Vec4{1, 2, 3, 4}
	rhs := &Vec4{5, 6, 7, 8}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Add(rhs)
	}
}

func TestMul(t *testing.T) {
	lhs := &Vec4{1, 2, 3, 4}
	var rhs float32 = 2.5
	lhs.Mul(rhs)
	if !reflect.DeepEqual(lhs, &Vec4{2.5, 5, 7.5, 10}) {
		t.Fatalf("Mul wrong result, got: %v", lhs)
	}
}

func BenchmarkMul(b *testing.B) {
	lhs := &Vec4{1, 2, 3, 4}
	var rhs float32 = 2.5
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		lhs.Mul(rhs)
	}
}
