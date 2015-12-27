package math32_test

import (
	"math"
	"testing"

	. "./"
)

func BenchmarkAbsFloat64(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		res := float32(math.Abs(float64(vf[n%count])))
		// work arround "evaluated but not used" error
		_ = res
	}
}

func BenchmarkAbsFloat32(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		res := Abs(vf[n%count])
		// work arround "evaluated but not used" error
		_ = res
	}
}

func BenchmarkCosFloat64(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		res := float32(math.Cos(float64(vf[n%count])))
		// work arround "evaluated but not used" error
		_ = res
	}
}

func BenchmarkCosFloat32(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		res := Cos(vf[n%count])
		// work arround "evaluated but not used" error
		_ = res
	}
}

func BenchmarkSinFloat64(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		res := float32(math.Sin(float64(vf[n%count])))
		// work arround "evaluated but not used" error
		_ = res
	}
}

func BenchmarkSinFloat32(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		res := Sin(vf[n%count])
		// work arround "evaluated but not used" error
		_ = res
	}
}

func BenchmarkSqrtFloat64(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		res := float32(math.Sqrt(float64(vf[n%count])))
		// work arround "evaluated but not used" error
		_ = res
	}
}

func BenchmarkSqrtFloat32(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		res := Sqrt(vf[n%count])
		// work arround "evaluated but not used" error
		_ = res
	}
}

func BenchmarkTanFloat64(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		res := float32(math.Tan(float64(vf[n%count])))
		// work arround "evaluated but not used" error
		_ = res
	}
}

func BenchmarkTanFloat32(b *testing.B) {
	count := len(vf)

	for n := 0; n < b.N; n++ {
		res := Tan(vf[n%count])
		// work arround "evaluated but not used" error
		_ = res
	}
}
