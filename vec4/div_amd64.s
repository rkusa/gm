#include "textflag.h"

// func divSIMD(lhs *Vec4, rhs float32)
TEXT ·divSIMD(SB),NOSPLIT,$0
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
