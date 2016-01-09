#include "textflag.h"

// func addSIMD(lhs, rhs *Vec4)
TEXT ·addSIMD(SB),NOSPLIT,$0
  // load pointers into registers
  MOVL lhs+0(FP), AX
  MOVL lhs+4(FP), BX

  // move unaligned vectors into SEE registers
  MOVUPS  (AX), X0
  MOVUPS  (BX), X1

  // add vector elements
  ADDPS   X1, X0

  // save result back into first vector
  MOVUPS X0, (AX)
  RET
