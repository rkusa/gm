// This file contains a subset of functions of the std
// math library from Go, but converted from float64 to float32.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32_test

import (
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
var fabs = []float32{
	4.9790119248836735e+00,
	7.7388724745781045e+00,
	2.7688005719200159e-01,
	5.0106036182710749e+00,
	9.6362937071984173e+00,
	2.9263772392439646e+00,
	5.2290834314593066e+00,
	2.7279399104360102e+00,
	1.8253080916808550e+00,
	8.6859247685756013e+00,
}
var sqrt = []float32{
	2.2313699855653004178179799e+00,
	2.7818829105618685382239619e+00,
	5.2619393351308207940064676e-01,
	2.2384377203502809905444337e+00,
	3.104237975937876203857968e+00,
	1.7106657465582673083304144e+00,
	2.2867189460131345235538447e+00,
	1.6516476149988743582497364e+00,
	1.3510396309834566963559155e+00,
	2.9471892592823585310668477e+00,
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

var vffabsSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var fabsSC = []float32{
	Inf(1),
	0,
	0,
	Inf(1),
	NaN(),
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

var vfsqrtSC = []float32{
	Inf(-1),
	-Pi,
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var sqrtSC = []float32{
	NaN(),
	NaN(),
	Copysign(0, -1),
	0,
	Inf(1),
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

func TestAbs(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Abs(vf[i]); fabs[i] != f {
			t.Errorf("Abs(%g) = %g, want %g", vf[i], f, fabs[i])
		}
	}
	for i := 0; i < len(vffabsSC); i++ {
		if f := Abs(vffabsSC[i]); !alike(fabsSC[i], f) {
			t.Errorf("Abs(%g) = %g, want %g", vffabsSC[i], f, fabsSC[i])
		}
	}
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

func TestSqrt(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		// a := Abs(vf[i])
		// if f := SqrtGo(a); sqrt[i] != f {
		// 	t.Errorf("SqrtGo(%g) = %g, want %g", a, f, sqrt[i])
		// }
		a := Abs(vf[i])
		if f := Sqrt(a); sqrt[i] != f {
			t.Errorf("Sqrt(%g) = %g, want %g", a, f, sqrt[i])
		}
	}
	for i := 0; i < len(vfsqrtSC); i++ {
		// if f := SqrtGo(vfsqrtSC[i]); !alike(sqrtSC[i], f) {
		// 	t.Errorf("SqrtGo(%g) = %g, want %g", vfsqrtSC[i], f, sqrtSC[i])
		// }
		if f := Sqrt(vfsqrtSC[i]); !alike(sqrtSC[i], f) {
			t.Errorf("Sqrt(%g) = %g, want %g", vfsqrtSC[i], f, sqrtSC[i])
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
