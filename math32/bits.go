// This file contains a subset of functions of the std
// math library from Go, but converted from float64 to float32.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import "math"

const (
	uvnan    = 0x7FC00001
	uvinf    = 0x7F800000
	uvneginf = 0xFF800000
)

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf(sign int) float32 {
	var v uint32
	if sign >= 0 {
		v = uvinf
	} else {
		v = uvneginf
	}
	return math.Float32frombits(v)
}

// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN() float32 { return math.Float32frombits(uvnan) }

// IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
func IsNaN(f float32) (is bool) {
	return f != f
}

// IsInf reports whether f is an infinity, according to sign.
// If sign > 0, IsInf reports whether f is positive infinity.
// If sign < 0, IsInf reports whether f is negative infinity.
// If sign == 0, IsInf reports whether f is either infinity.
func IsInf(f float32, sign int) bool {
	return sign >= 0 && f > math.MaxFloat32 || sign <= 0 && f < -math.MaxFloat32
}
