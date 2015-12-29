#include "textflag.h"

// func lenSIMD(lhs *Vec4) float32
TEXT ·lenSIMD(SB),NOSPLIT,$0
  // load pointer into register
  MOVQ    lhs+0(FP), R8

  // move vector into SEE register
  MOVUPS  (R8), X0

  // multipliy with itself
  MULPS   X0, X0

  // copy two high values to low values
  MOVHLPS X0, X1

  // add two low values
  ADDPS   X1, X0

  // copy low values reverse
  MOVUPS  X0, X1
  SHUFPS  $0x01, X1, X1
  ADDSS   X1, X0

  // sqrt resulting sum
  SQRTSS  X0, X0

  // return result
  MOVSS   X0, ret+8(FP)
  RET
