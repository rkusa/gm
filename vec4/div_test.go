package vec4

import (
	"reflect"
	"testing"
)

func TestDivGo(t *testing.T) {
	testDiv(t, div)
}

func TestDivSIMD(t *testing.T) {
	testDiv(t, divSIMD)
}

func TestDiv(t *testing.T) {
	testDiv(t, func(lhs *Vec4, rhs float32) {
		lhs.Div(rhs)
	})
}

func testDiv(t *testing.T, div func(lhs *Vec4, rhs float32)) {
	lhs := &Vec4{1, 2, 3, 4}
	var rhs float32 = 2
	div(lhs, rhs)

	if !reflect.DeepEqual(lhs, &Vec4{.5, 1, 1.5, 2}) {
		t.Fatalf("Div wrong result, got: %v", lhs)
	}
}

func BenchmarkDivGo(b *testing.B) {
	benchmarkDiv(b, div)
}

func BenchmarkDivSIMD(b *testing.B) {
	benchmarkDiv(b, divSIMD)
}

func benchmarkDiv(b *testing.B, div func(lhs *Vec4, rhs float32)) {
	lhs := &Vec4{1, 2, 3, 4}
	var rhs float32 = 2.5
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		div(lhs, rhs)
	}
}
