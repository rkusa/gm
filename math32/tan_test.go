// This file contains a subset of functions of the std
// math library from Go, but converted from float64 to float32.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32_test

import (
	"math"
	"testing"

	. "./"
)

var vf = []float32{
	4.9790119248836735e+00,
	7.7388724745781045e+00,
	-2.7688005719200159e-01,
	-5.0106036182710749e+00,
	9.6362937071984173e+00,
	2.9263772392439646e+00,
	5.2290834314593066e+00,
	2.7279399104360102e+00,
	1.8253080916808550e+00,
	-8.6859247685756013e+00,
}
var copysign = []float32{
	-4.9790119248836735e+00,
	-7.7388724745781045e+00,
	-2.7688005719200159e-01,
	-5.0106036182710749e+00,
	-9.6362937071984173e+00,
	-2.9263772392439646e+00,
	-5.2290834314593066e+00,
	-2.7279399104360102e+00,
	-1.8253080916808550e+00,
	-8.6859247685756013e+00,
}
var tan = []float32{
	-3.661316565040227801781974e+00,
	8.64900232648597589369854e+00,
	-2.8417941955033612725238097e-01,
	3.253290185974728640827156e+00,
	2.147275640380293804770778e-01,
	-2.18600910711067004921551e-01,
	-1.760002817872367935518928e+00,
	-4.389808914752818126249079e-01,
	-3.843885560201130679995041e+00,
	9.10988793377685105753416e-01,
}

var vfcopysignSC = []float32{
	Inf(-1),
	Inf(1),
	NaN(),
}
var copysignSC = []float32{
	Inf(-1),
	Inf(-1),
	NaN(),
}
var vfsinSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var sinSC = []float32{
	NaN(),
	Copysign(0, -1),
	0,
	NaN(),
	NaN(),
}

func tolerance(a, b, e float32) bool {
	d := a - b
	if d < 0 {
		d = -d
	}

	// note: b is correct (expected) value, a is actual value.
	// make error tolerance a fraction of b, not a.
	if b != 0 {
		e = e * b
		if e < 0 {
			e = -e
		}
	}
	return d < e
}
func veryclose(a, b float32) bool { return tolerance(a, b, 4e-6) }
func alike(a, b float32) bool {
	switch {
	case IsNaN(a) && IsNaN(b):
		return true
	case a == b:
		return Signbit(a) == Signbit(b)
	}
	return false
}

func TestCopysign(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Copysign(vf[i], -1); copysign[i] != f {
			t.Errorf("Copysign(%g, -1) = %g, want %g", vf[i], f, copysign[i])
		}
	}
	for i := 0; i < len(vf); i++ {
		if f := Copysign(vf[i], 1); -copysign[i] != f {
			t.Errorf("Copysign(%g, 1) = %g, want %g", vf[i], f, -copysign[i])
		}
	}
	for i := 0; i < len(vfcopysignSC); i++ {
		if f := Copysign(vfcopysignSC[i], -1); !alike(copysignSC[i], f) {
			t.Errorf("Copysign(%g, -1) = %g, want %g", vfcopysignSC[i], f, copysignSC[i])
		}
	}
}

func TestTan(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Tan(vf[i]); !veryclose(tan[i], f) {
			t.Errorf("Tan(%g) = %g, want %g", vf[i], f, tan[i])
		}
	}
	// same special cases as Sin
	for i := 0; i < len(vfsinSC); i++ {
		if f := Tan(vfsinSC[i]); !alike(sinSC[i], f) {
			t.Errorf("Tan(%g) = %g, want %g", vfsinSC[i], f, sinSC[i])
		}
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
