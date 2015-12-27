#include "textflag.h"

// func addVec4SIMD(lhs, rhs *Vec4)
TEXT ·addVec4SIMD(SB),NOSPLIT,$0
  // load pointers into registers
  MOVQ lhs+0(FP), R8
  MOVQ rhs+8(FP), R9

  // move unaligned vectors into SEE registers
  MOVUPS (R8), X0
  MOVUPS (R9), X1

  // add vector elements
  ADDPS X1, X0

  // save result back into first vector
  MOVUPS X0, (R8)
  RET

// func divVec4SIMD(lhs *Vec4, rhs float32)
TEXT ·divVec4SIMD(SB),NOSPLIT,$0
  // load vector
  MOVQ lhs+0(FP), R8
  MOVUPS (R8), X0

  // load scalar
  MOVSS rhs+8(FP), X1
  // broadcast the lower element to all four fields
  SHUFPS $0x00, X1, X1

  // division
  DIVPS X1, X0

  // save result back into vector
  MOVUPS X0, (R8)
  RET

// func lenVec4SIMD(lhs *Vec4) float32
TEXT ·lenVec4SIMD(SB),NOSPLIT,$0
  // load pointer into register
  MOVQ lhs+0(FP), R8

  // move vector into SEE register
  MOVUPS (R8), X0

  // multipliy with itself
  MULPS X0, X0

  // copy two high values to low values
  MOVHLPS X0, X1

  // add two low values
  ADDPS X1, X0

  // copy low values reverse
  MOVUPS X0, X1
  SHUFPS $0x01, X1, X1
  ADDSS X1, X0

  // sqrt resulting sum
  SQRTSS X0, X0

  // return result
  MOVSS X0, ret+8(FP)
  RET

// func mulVec4SIMD(lhs *Vec4, rhs float32)
TEXT ·mulVec4SIMD(SB),NOSPLIT,$0
  // load vector
  MOVQ lhs+0(FP), R8
  MOVUPS (R8), X0

  // load scalar
  MOVSS rhs+8(FP), X1
  // broadcast the lower element to all four fields
  SHUFPS $0x00, X1, X1

  // multiply
  MULPS X1, X0

  // save result back into vector
  MOVUPS X0, (R8)
  RET

// func subVec4SIMD(lhs, rhs *Vec4)
TEXT ·subVec4SIMD(SB),NOSPLIT,$0
  // load pointers into registers
  MOVQ lhs+0(FP), R8
  MOVQ rhs+8(FP), R9

  // move unaligned vectors into SEE registers
  MOVUPS (R8), X0
  MOVUPS (R9), X1

  // substract vector elements
  SUBPS X1, X0

  // save result back into first vector
  MOVUPS X0, (R8)
  RET

