package math32_test

import (
	"math"
	"testing"

	. "./"
)

// benchmark results will be stored in this package level variable to avoid
// compiler optimisations eliminating the benchmarked functions
var result float32

func BenchmarkAbsFloat64(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		result = float32(math.Abs(float64(vf[n%count])))
	}
}

func BenchmarkAbsFloat32(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		result = Abs(vf[n%count])
	}
}

func BenchmarkCosFloat64(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		result = float32(math.Cos(float64(vf[n%count])))
	}
}

func BenchmarkCosFloat32(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		result = Cos(vf[n%count])
	}
}

func BenchmarkSinFloat64(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		result = float32(math.Sin(float64(vf[n%count])))
	}
}

func BenchmarkSinFloat32(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		result = Sin(vf[n%count])
	}
}

func BenchmarkSqrtFloat64(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		result = float32(math.Sqrt(float64(vf[n%count])))
	}
}

func BenchmarkSqrtFloat32(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		result = Sqrt(vf[n%count])
	}
}

func BenchmarkTanFloat64(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		result = float32(math.Tan(float64(vf[n%count])))
	}
}

func BenchmarkTanFloat32(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		result = Tan(vf[n%count])
	}
}
