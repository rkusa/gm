// This file contains a subset of functions of the std
// math library from Go, but converted from float64 to float32.

// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"


// func Sqrt(x float32) float32
TEXT ·Sqrt(SB),NOSPLIT,$0
  FMOVF   x+0(FP),F0
  FSQRT
  FMOVFP  F0,ret+4(FP)
  RET

