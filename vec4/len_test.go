package vec4

import "testing"

func TestLenGo(t *testing.T) {
	testLen(t, len)
}

func TestLenSIMD(t *testing.T) {
	testLen(t, lenSIMD)
}

func TestLen(t *testing.T) {
	testLen(t, (*Vec4).Len)
}

func testLen(t *testing.T, len func(lhs *Vec4) float32) {
	lhs := &Vec4{1, 2, 3, 4}

	l := len(lhs)
	if l != 5.477226 {
		t.Fatalf("Len wrong result, got: %v, %v", l, lhs)
	}
}

func BenchmarkLenGo(b *testing.B) {
	benchmarkLen(b, len)
}

func BenchmarkLenSIMD(b *testing.B) {
	benchmarkLen(b, lenSIMD)
}

func benchmarkLen(b *testing.B, len func(lhs *Vec4) float32) {
	lhs := &Vec4{1, 2, 3, 4}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		len(lhs)
	}
}
