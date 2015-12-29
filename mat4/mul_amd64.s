#include "textflag.h"

// func mulSIMD(out, lhs, rhs *Mat4)
TEXT ·mulSIMD(SB),NOSPLIT,$0
  // load pointers into registers
  MOVQ    lhs+8(FP), R8 // lhs
  MOVQ    rhs+16(FP), R9 // rhs
  MOVQ    rhs+0(FP), R10 // out

  // load lhs into SSE registers
  MOVUPS  (R8), X0   // row 1
  MOVUPS  16(R8), X1 // row 2
  MOVUPS  32(R8), X2 // row 3
  MOVUPS  48(R8), X3 // row 4

  ///  rhs row 1

  // mul lhs[0] with rhs[0][0]
  MOVUPS  (R9), X4       // load rhs row 1
  SHUFPS  $0x00, X4, X4  // broadcast 1. val
  MULPS   X0, X4          // mul with lhs row 1

  // mul lhs[1] with rhs[0][1]
  MOVUPS  (R9), X5       // load rhs row 1
  SHUFPS  $0x55, X5, X5  // broadcast 2. val
  MULPS   X1, X5          // mul with lhs row 2

  // mul lhs[2] with rhs[0][2]
  MOVUPS  (R9), X6       // load rhs row 1
  SHUFPS  $0xAA, X6, X6  // broadcast 3. val
  MULPS   X2, X6          // mul with lhs row 3

  // mul lhs[3] with rhs[0][3]
  MOVUPS  (R9), X7       // load rhs row 1
  SHUFPS  $0xFF, X7, X7  // broadcast 4. val
  MULPS   X3, X7          // mul with lhs row 4

  // add results
  ADDPS   X5, X4
  ADDPS   X6, X4
  ADDPS   X7, X4

  // save result row
  MOVUPS  X4, (R10)

  ///  rhs row 2

  // mul lhs[0] with rhs[1][0]
  MOVUPS  16(R9), X4       // load rhs row 2
  SHUFPS  $0x00, X4, X4  // broadcast 1. val
  MULPS   X0, X4          // mul with lhs row 1

  // mul lhs[1] with rhs[1][1]
  MOVUPS  16(R9), X5       // load rhs row 2
  SHUFPS  $0x55, X5, X5  // broadcast 2. val
  MULPS   X1, X5          // mul with lhs row 2

  // mul lhs[2] with rhs[1][2]
  MOVUPS  16(R9), X6       // load rhs row 2
  SHUFPS  $0xAA, X6, X6  // broadcast 3. val
  MULPS   X2, X6          // mul with lhs row 3

  // mul lhs[3] with rhs[1][3]
  MOVUPS  16(R9), X7       // load rhs row 2
  SHUFPS  $0xFF, X7, X7  // broadcast 4. val
  MULPS   X3, X7          // mul with lhs row 4

  // add results
  ADDPS   X5, X4
  ADDPS   X6, X4
  ADDPS   X7, X4

  // save result row
  MOVUPS  X4, 16(R10)

  ///  rhs row 3

  // mul lhs[0] with rhs[2][0]
  MOVUPS  32(R9), X4     // load rhs row 3
  SHUFPS  $0x00, X4, X4  // broadcast 1. val
  MULPS   X0, X4          // mul with lhs row 1

  // mul lhs[1] with rhs[2][1]
  MOVUPS  32(R9), X5     // load rhs row 3
  SHUFPS  $0x55, X5, X5  // broadcast 2. val
  MULPS   X1, X5          // mul with lhs row 2

  // mul lhs[2] with rhs[2][2]
  MOVUPS  32(R9), X6     // load rhs row 3
  SHUFPS  $0xAA, X6, X6  // broadcast 3. val
  MULPS   X2, X6          // mul with lhs row 3

  // mul lhs[3] with rhs[2][3]
  MOVUPS  32(R9), X7     // load rhs row 3
  SHUFPS  $0xFF, X7, X7  // broadcast 4. val
  MULPS   X3, X7          // mul with lhs row 4

  // add results
  ADDPS   X5, X4
  ADDPS   X6, X4
  ADDPS   X7, X4

  // save result row
  MOVUPS  X4, 32(R10)

  ///  rhs row 4

  // mul lhs[0] with rhs[3][0]
  MOVUPS  48(R9), X4     // load rhs row 3
  SHUFPS  $0x00, X4, X4  // broadcast 1. val
  MULPS   X0, X4          // mul with lhs row 1

  // mul lhs[1] with rhs[3][1]
  MOVUPS  48(R9), X5     // load rhs row 3
  SHUFPS  $0x55, X5, X5  // broadcast 2. val
  MULPS   X1, X5          // mul with lhs row 2

  // mul lhs[2] with rhs[3][2]
  MOVUPS  48(R9), X6     // load rhs row 3
  SHUFPS  $0xAA, X6, X6  // broadcast 3. val
  MULPS   X2, X6          // mul with lhs row 3

  // mul lhs[3] with rhs[3][3]
  MOVUPS  48(R9), X7     // load rhs row 3
  SHUFPS  $0xFF, X7, X7  // broadcast 4. val
  MULPS   X3, X7          // mul with lhs row 4

  // add results
  ADDPS   X5, X4
  ADDPS   X6, X4
  ADDPS   X7, X4

  // save result row
  MOVUPS  X4, 48(R10)

  /// done
  RET
