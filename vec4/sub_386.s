#include "textflag.h"

// func subSIMD(lhs, rhs *Vec4)
TEXT ·subSIMD(SB),NOSPLIT,$0
  // load pointers into registers
  MOVL    lhs+0(FP), AX
  MOVL    rhs+4(FP), BX

  // move unaligned vectors into SEE registers
  MOVUPS  (AX), X0
  MOVUPS  (BX), X1

  // substract vector elements
  SUBPS   X1, X0

  // save result back into first vector
  MOVUPS  X0, (AX)
  RET
