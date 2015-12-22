// func add(lhs, rhs *Vec4)
TEXT ·add(SB),NOSPLIT,$0
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

// func mul(lhs *Vec4, rhs float32)
TEXT ·mul(SB),NOSPLIT,$0
  // load vector
  MOVQ lhs+0(FP), R8
  MOVUPS (R8), X0

  // load scalar
  MOVLPS rhs+8(FP), X1
  // broadcast the lower element to all four fields
  SHUFPS $0x00, X1, X1

  // multiply
  MULPS X1, X0

  // save result back into vector
  MOVUPS X0, (R8)
  RET
