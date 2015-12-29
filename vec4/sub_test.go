package vec4

import (
	"reflect"
	"testing"
)

func TestSubGo(t *testing.T) {
	testSub(t, sub)
}

func TestSubSIMD(t *testing.T) {
	testSub(t, subSIMD)
}

func TestSub(t *testing.T) {
	testSub(t, func(lhs, rhs *Vec4) {
		lhs.Sub(rhs)
	})
}

func testSub(t *testing.T, sub func(lhs, rhs *Vec4)) {
	lhs := &Vec4{1, 2, 3, 4}
	rhs := &Vec4{8, 7, 6, 5}

	sub(lhs, rhs)
	if !reflect.DeepEqual(lhs, &Vec4{-7, -5, -3, -1}) {
		t.Fatalf("Sub wrong result, got: %v", lhs)
	}

	// test sub itself
	lhs = &Vec4{1, 2, 3, 4}

	sub(lhs, lhs)
	if !reflect.DeepEqual(lhs, &Vec4{0, 0, 0, 0}) {
		t.Fatalf("Sub itself wrong result, got: %v", lhs)
	}
}

func BenchmarkSubGo(b *testing.B) {
	benchmarkSub(b, sub)
}

func BenchmarkSubSIMD(b *testing.B) {
	benchmarkSub(b, subSIMD)
}

func benchmarkSub(b *testing.B, sub func(lhs, rhs *Vec4)) {
	lhs := &Vec4{1, 2, 3, 4}
	rhs := &Vec4{5, 6, 7, 8}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		sub(lhs, rhs)
	}
}
