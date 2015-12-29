package mat4

import (
	"reflect"
	"testing"
)

func TestInvertGo(t *testing.T) {
	testInvert(t, invert)
}

func TestInvertSIMD(t *testing.T) {
	testInvert(t, invertSIMD)
}

func TestInvert(t *testing.T) {
	testInvert(t, func(lhs *Mat4) bool {
		return lhs.Invert() != nil
	})
}

func testInvert(t *testing.T, invert func(lhs *Mat4) bool) {
	lhs := &Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 0,
	}
	clone := lhs.Clone()

	if invert(lhs) {
		t.Fatalf("Invert wrong result, does not break on 0 row")
	}

	if !reflect.DeepEqual(lhs, clone) {
		t.Fatalf("Invert should not mutate on 0 row")
	}

	lhs[15] = 1
	expectation := &Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		-1, -2, -3, 1,
	}

	if !invert(lhs) {
		t.Fatalf("Invert should return true if there is no 0 row")
	}

	if lhs == nil || !close(lhs, expectation) {
		t.Fatalf("Invert wrong result, got: %v", lhs)
	}
}

func BenchmarkInvertGo(b *testing.B) {
	benchmarkInvert(b, invert)
}

func BenchmarkInvertSIMD(b *testing.B) {
	benchmarkInvert(b, invertSIMD)
}

func benchmarkInvert(b *testing.B, invert func(lhs *Mat4) bool) {
	lhs := Identity()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		invert(lhs)
	}
}
