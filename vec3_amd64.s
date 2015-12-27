#include "textflag.h"

// func lenVec3SIMD(lhs *Vec3) float32
TEXT ·lenVec3SIMD(SB),NOSPLIT,$0
  // load pointer into register
  MOVQ lhs+0(FP), R8

  // move vector into SEE register
  MOVUPS (R8), X0

  // multipliy with itself
  MULPS X0, X0

  // move lower value to X1
  MOVSS X0, X1

  // two times: shift right & add lower value to X1
  SHUFPS $0x39, X0, X0
  ADDSS X0, X1
  SHUFPS $0x39, X0, X0
  ADDSS X0, X1

  // sqrt resulting sum
  SQRTSS X1, X1

  // return result
  MOVSS X1, ret+8(FP)
  RET

// func subVec3SIMD(lhs, rhs *Vec3)
TEXT ·subVec3SIMD(SB),NOSPLIT,$0
  // load pointers into registers
  MOVQ lhs+0(FP), R8
  MOVQ rhs+8(FP), R9

  // move unaligned vectors into SEE registers
  MOVUPS (R8), X0
  MOVUPS (R9), X1

  // add vector elements
  SUBPS X1, X0

  // save result back into first vector
  MOVSS X0, (R8)        // move lower value
  SHUFPS $0x39, X0, X0  // shift right
  MOVSD X0, +4(R8)      // move lower two values
  RET

// func divVec3SIMD(lhs *Vec3, rhs float32)
TEXT ·divVec3SIMD(SB),NOSPLIT,$0
  // load vector
  MOVQ lhs+0(FP), R8
  MOVUPS (R8), X0

  // load scalar
  MOVSS rhs+8(FP), X1
  // broadcast the lower element to all four fields
  SHUFPS $0x00, X1, X1

  // multiply
  DIVPS X1, X0

  // save result back into vector
  MOVSS X0, (R8)        // move lower value
  SHUFPS $0x39, X0, X0  // shift right
  MOVSD X0, +4(R8)      // move lower two values
  RET
