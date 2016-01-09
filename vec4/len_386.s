#include "textflag.h"

// func lenSIMD(lhs *Vec4) float32
TEXT ·lenSIMD(SB),NOSPLIT,$0
  // load pointer into register
  MOVL    lhs+0(FP), AX

  // move vector into SEE register
  MOVUPS  (AX), X0

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
  MOVSS   X0, ret+4(FP)
  RET
