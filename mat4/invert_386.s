#include "textflag.h"

// func invertSIMD(lhs *Mat4)
TEXT ·invertSIMD386(SB), NOSPLIT, $0
  // implementation based on:
  // http://www.intel.com/design/pentiumiii/sml/245043.htm

  // X4 ... det
  // X5 ... tmp
  // X6 ... tmp2

  // X0 ... row0
  // X1 ... row1
  // X2 ... row2
  // X3 ... row3

  // load pointers into registers
  //   (AX) ... lhs0
  // 16(AX) ... lhs1
  // 32(AX) ... lhs2
  // 48(AX) ... lhs3
  MOVL lhs+0(FP), AX

  // temporary values
  MOVL lhs+8(FP), BX

  // tmp1
  MOVUPS (AX), X5          // lhs0 -> tmp
  SHUFPS $0x44, 16(AX), X5 // mask: 01 00 01 00

  // row1
  MOVUPS 32(AX), X1        // lhs2 -> row1
  SHUFPS $0x44, 48(AX), X1 // mask: 01 00 01 00

  // row0
  MOVUPS X5, X0       // row1 -> row0
  SHUFPS $0x88, X1, X0 // mask: 10 00 10 00

  // row1
  SHUFPS $0xDD, X5, X1 // mask: 11 01 11 01

  // tmp1
  MOVUPS (AX), X5          // lhs0 -> tmp1
  SHUFPS $0xEE, 16(AX), X5 // mask: 11 10 11 10

  // row3
  MOVUPS 32(AX), X3        // lhs2 -> row3
  SHUFPS $0xEE, 48(AX), X3 // mask: 11 10 11 10

  // row2
  MOVUPS X5, X2        // row3 -> row2
  SHUFPS $0x88, X3, X2 // mask: 10 00 10 00

  // row3
  SHUFPS $0xDD, X5, X3

  // ---------------------------------------

  // save rows for later use
  MOVUPS X0, (BX)
  MOVUPS X1, 16(BX)
  MOVUPS X2, 32(BX)
  MOVUPS X3, 48(BX)

  // ---------------------------------------

  // from now on>
  // X0 ... minor0
  // X1 ... minor1
  // X2 ... minor2
  // X3 ... minor3

  // tmp1
  MOVUPS 32(BX), X5        // row2 -> tmp
  MULPS  48(BX), X5        // row3 * row2 (tmp)
  SHUFPS $0xB1, X5, X5

  // minor0
  MOVUPS 16(BX), X0  // row1 -> minor0
  MULPS  X5, X0 // tmp * row1 (minor0)

  // minor1
  MOVUPS (BX), X1  // row0 -> minor1
  MULPS  X5, X1 // tmp * row0 (minor1)

  // tmp1
  SHUFPS $0x4E, X5, X5

  // minor0
  MOVUPS X0, X6 // temp minor0
  MOVUPS 16(BX), X0  // row1 -> minor0
  MULPS  X5, X0 // tmp * row1 (minor0)
  SUBPS  X6, X0 // minor0 - tmp

  // minor1
  MOVUPS X1, X6       // temp minor1
  MOVUPS (BX), X1        // row0 -> minor1
  MULPS  X5, X1       // tmp * row0 (minor1)
  SUBPS  X6, X1       // minor1 - tmp
  SHUFPS $0x4E, X1, X1

  // ---------------------------------------

  // tmp1
  MOVUPS 16(BX), X5         // row1 -> tmp
  MULPS  32(BX), X5        // row2 * row1 (tmp)
  SHUFPS $0xB1, X5, X5

  // minor0
  MOVUPS X0, X6 // temp minor0
  MOVUPS 48(BX), X0 // row3 -> minor0
  MULPS  X5, X0 // tmp * row3 (minor0)
  ADDPS  X6, X0 // minor0 - tmp

  // minor3
  MOVUPS (BX), X3  // row0 -> minor3
  MULPS  X5, X3 // tmp * row0 (minor3)

  // tmp
  SHUFPS $0x4E, X5, X5

  // minor0
  MOVUPS 48(BX), X6 // row3 -> temporary
  MULPS  X5, X6 // tmp * row3 (temporary)
  SUBPS  X6, X0  // minor0 - temporary

  // minor3
  MOVUPS X3, X6       // minor3 -> temporary
  MOVUPS (BX), X3        // row0 -> minor3
  MULPS  X5, X3       // tmp * row0 (minor1)
  SUBPS  X6, X3       // minor1 - tmp
  SHUFPS $0x4E, X3, X3

  // ---------------------------------------

  // tmp
  MOVUPS 16(BX), X5         // row1 -> tmp
  SHUFPS $0x4E, X5, X5
  MULPS  48(BX), X5        // row3 * tmp
  SHUFPS $0xB1, X5, X5

  // row2
  MOVUPS 32(BX), X6
  SHUFPS $0x4E, X6, X6
  MOVUPS X6, 32(BX)

  // minor0
  MOVUPS 32(BX), X6 // row2 -> temporary
  MULPS  X5, X6 // tmp * row2 (temporary)
  ADDPS  X6, X0  // minor0 - temporary

  // minor2
  MOVUPS (BX), X2  // row2 -> minor2
  MULPS  X5, X2 // tmp * minor2

  // tmp
  SHUFPS $0x4E, X5, X5

  // minor0
  MOVUPS 32(BX), X6 // row2 -> temporary
  MULPS  X5, X6 // tmp * row2 (temporary)
  SUBPS  X6, X0  // minor0 - temporary

  // minor2
  MOVUPS X2, X6       // minor2 -> temporary
  MOVUPS (BX), X2        // row0 -> minor2
  MULPS  X5, X2       // tmp * row0 (minor1)
  SUBPS  X6, X2       // minor1 - tmp
  SHUFPS $0x4E, X2, X2

  // ---------------------------------------

  // tmp
  MOVUPS (BX), X5         // row0 -> tmp
  MULPS  16(BX), X5         // row1 * row2 (tmp)
  SHUFPS $0xB1, X5, X5

  // minor2
  MOVUPS X2, X6 // minor2 -> temporary
  MOVUPS 48(BX), X2 // row3 -> minor2
  MULPS  X5, X2 // tmp * row3 (minor2)
  ADDPS  X6, X2 // add

  // minor3
  MOVUPS X3, X6 // minor3 -> temporary
  MOVUPS 32(BX), X3 // row2 -> minor3
  MULPS  X5, X3 // tmp * row2 (minor3)
  SUBPS  X6, X3 // sub

  // tmp
  SHUFPS $0x4E, X5, X5

  // minor2
  MOVUPS X2, X6 // minor2 -> temporary
  MOVUPS 48(BX), X2 // row3 -> minor2
  MULPS  X5, X2 // tmp * row3 (minor2)
  SUBPS  X6, X2 // sub

  // minor3
  MOVUPS 32(BX), X6 // row2 -> temporary
  MULPS  X5, X6 // tmp * row2 (temporary)
  SUBPS  X6, X3  // sub

  // ---------------------------------------

  // tmp
  MOVUPS (BX), X5         // row0 -> tmp
  MULPS  48(BX), X5        // row3 * row0 (tmp)
  SHUFPS $0xB1, X5, X5

  // minor1
  MOVUPS 32(BX), X6 // row2 -> temporary
  MULPS  X5, X6 // tmp * row2 (temporary)
  SUBPS  X6, X1  // sub

  // minor2
  MOVUPS X2, X6 // minor2 -> temporary
  MOVUPS 16(BX), X2  // row1 -> minor2
  MULPS  X5, X2 // tmp * row3 (minor2)
  ADDPS  X6, X2 // add

  // tmp
  SHUFPS $0x4E, X5, X5

  // minor1
  MOVUPS X1, X6 // minor2 -> temporary
  MOVUPS 32(BX), X1 // row2 -> minor1
  MULPS  X5, X1 // tmp * row3 (minor1)
  ADDPS  X6, X1 // add

  // minor2
  MOVUPS 16(BX), X6  // row2 -> temporary
  MULPS  X5, X6 // tmp * row2 (temporary)
  SUBPS  X6, X2  // sub

  // ---------------------------------------

  // tmp
  MOVUPS (BX), X5         // row0 -> tmp
  MULPS  32(BX), X5        // row2 * row0 (tmp)
  SHUFPS $0xB1, X5, X5

  // minor1
  MOVUPS X1, X6 // minor1 -> temporary
  MOVUPS 48(BX), X1 // row3 -> minor1
  MULPS  X5, X1 // tmp * row3 (minor1)
  ADDPS  X6, X1 // sub

  // minor3
  MOVUPS 16(BX), X6  // row1 -> temporary
  MULPS  X5, X6 // tmp * row1 (temporary)
  SUBPS  X6, X3  // sub

  // tmp
  SHUFPS $0x4E, X5, X5

  // minor1
  MOVUPS 48(BX), X6 // row3 -> temporary
  MULPS  X5, X6 // tmp * row3 (temporary)
  SUBPS  X6, X1  // sub

  // minor3
  MOVUPS X3, X6 // minor3 -> temporary
  MOVUPS 16(BX), X3  // row1 -> minor3
  MULPS  X5, X3 // tmp * row1 (minor3)
  ADDPS  X6, X3 // add

  // ---------------------------------------

  // det
  MOVUPS (BX), X4         // row0 -> det
  MULPS  X0, X4         // minor0 * row0 (det)
  MOVUPS X4, X6        // det -> temporary
  SHUFPS $0x4E, X6, X6
  ADDPS  X6, X4
  MOVUPS X4, X6        // det -> temporary
  SHUFPS $0xB1, X6, X6
  ADDSS  X6, X4

  // break on det == 0
  MOVL    $0, BX
  MOVL    BX, X6
  UCOMISS X6, X4
  JNZ     result        // jump if not zero
  MOVB    $0, ret+4(FP) // return false
  RET

result:

  // tmp
  MOVUPS X4, X5 // det -> tmp
  RCPSS  X5, X5 // reciprical

  // det
  MOVUPS X5, X6        // tmp -> temporary
  MULSS  X6, X4
  ADDSS  X5, X5
  SUBSS  X4, X5
  MOVUPS X5, X4
  SHUFPS $0x00, X4, X4

  // minor0
  MULPS  X4, X0
  MOVUPS X0, (AX) // write result

  // minor1
  MULPS  X4, X1
  MOVUPS X1, 16(AX) // write result

  // minor2
  MULPS  X4, X2
  MOVUPS X2, 32(AX) // write result

  // minor3
  MULPS  X4, X3
  MOVUPS X3, 48(AX) // write result

  // ---------------------------------------

  MOVB $1, CL
  MOVB CL, ret+4(FP) // return true
  RET
