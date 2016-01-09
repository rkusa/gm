#include "textflag.h"

// func mulSIMD(lhs *Vec4, rhs float32)
TEXT ·mulSIMD(SB),NOSPLIT,$0
  // load vector
  MOVL    lhs+0(FP), AX
  MOVUPS  (AX), X0

  // load scalar
  MOVSS   rhs+4(FP), X1
  // broadcast the lower element to all four fields
  SHUFPS  $0x00, X1, X1

  // multiply
  MULPS   X1, X0

  // save result back into vector
  MOVUPS  X0, (AX)
  RET
