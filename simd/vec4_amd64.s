// func add(lhs, rhs *Vec4)
TEXT ·add(SB),NOSPLIT,$0
  // load pointers into registers
  MOVQ lhs+0(FP), R8
  MOVQ rhs+8(FP), R9

  // move unaligned vectors into SEE registers
  MOVUPS (R8), X8
  MOVUPS (R9), X9

  // add vector elements
  ADDPS X9, X8

  // save result into first argument
  MOVUPS X8, (R8)

  // done
  RET
