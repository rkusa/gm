package vec4

import (
	"reflect"
	"testing"
)

func TestAddGo(t *testing.T) {
	testAdd(t, add)
}

func TestAddSIMD(t *testing.T) {
	testAdd(t, addSIMD)
}

func TestAdd(t *testing.T) {
	testAdd(t, func(lhs, rhs *Vec4) {
		lhs.Add(rhs)
	})
}

func testAdd(t *testing.T, add func(lhs, rhs *Vec4)) {
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

func BenchmarkAddGo(b *testing.B) {
	benchmarkAdd(b, add)
}

func BenchmarkAddSIMD(b *testing.B) {
	benchmarkAdd(b, addSIMD)
}

func benchmarkAdd(b *testing.B, add func(lhs, rhs *Vec4)) {
	lhs := &Vec4{1, 2, 3, 4}
	rhs := &Vec4{5, 6, 7, 8}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		add(lhs, rhs)
	}
}
