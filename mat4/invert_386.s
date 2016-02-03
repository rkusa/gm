#include "textflag.h"

// func invertSIMD(lhs *Mat4)
TEXT ·invertSIMD(SB), NOSPLIT, $0
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
  //   (R8) ... lhs0
  // 16(R8) ... lhs1
  // 32(R8) ... lhs2
  // 48(R8) ... lhs3
  MOVQ lhs+0(FP), R8

  // temporary values
  MOVQ lhs+8(FP), R9

  // tmp1
  MOVUPS (R8), X5          // lhs0 -> tmp
  SHUFPS $0x44, 16(R8), X5 // mask: 01 00 01 00

  // row1
  MOVUPS 32(R8), X1        // lhs2 -> row1
  SHUFPS $0x44, 48(R8), X1 // mask: 01 00 01 00

  // row0
  MOVUPS X5, X0       // row1 -> row0
  SHUFPS $0x88, X1, X0 // mask: 10 00 10 00

  // row1
  SHUFPS $0xDD, X5, X1 // mask: 11 01 11 01

  // tmp1
  MOVUPS (R8), X5          // lhs0 -> tmp1
  SHUFPS $0xEE, 16(R8), X5 // mask: 11 10 11 10

  // row3
  MOVUPS 32(R8), X3        // lhs2 -> row3
  SHUFPS $0xEE, 48(R8), X3 // mask: 11 10 11 10

  // row2
  MOVUPS X5, X2        // row3 -> row2
  SHUFPS $0x88, X3, X2 // mask: 10 00 10 00

  // row3
  SHUFPS $0xDD, X5, X3

  // ---------------------------------------

  // save rows for later use
  MOVUPS X0, (R9)
  MOVUPS X1, 16(R9)
  MOVUPS X2, 32(R9)
  MOVUPS X3, 48(R9)

  // ---------------------------------------

  // from now on>
  // X0 ... minor0
  // X1 ... minor1
  // X2 ... minor2
  // X3 ... minor3

  // tmp1
  MOVUPS 32(R9), X5        // row2 -> tmp
  MULPS  48(R9), X5        // row3 * row2 (tmp)
  SHUFPS $0xB1, X5, X5

  // minor0
  MOVUPS 16(R9), X0  // row1 -> minor0
  MULPS  X5, X0 // tmp * row1 (minor0)

  // minor1
  MOVUPS (R9), X1  // row0 -> minor1
  MULPS  X5, X1 // tmp * row0 (minor1)

  // tmp1
  SHUFPS $0x4E, X5, X5

  // minor0
  MOVUPS X0, X6 // temp minor0
  MOVUPS 16(R9), X0  // row1 -> minor0
  MULPS  X5, X0 // tmp * row1 (minor0)
  SUBPS  X6, X0 // minor0 - tmp

  // minor1
  MOVUPS X1, X6       // temp minor1
  MOVUPS (R9), X1        // row0 -> minor1
  MULPS  X5, X1       // tmp * row0 (minor1)
  SUBPS  X6, X1       // minor1 - tmp
  SHUFPS $0x4E, X1, X1

  // ---------------------------------------

  // tmp1
  MOVUPS 16(R9), X5         // row1 -> tmp
  MULPS  32(R9), X5        // row2 * row1 (tmp)
  SHUFPS $0xB1, X5, X5

  // minor0
  MOVUPS X0, X6 // temp minor0
  MOVUPS 48(R9), X0 // row3 -> minor0
  MULPS  X5, X0 // tmp * row3 (minor0)
  ADDPS  X6, X0 // minor0 - tmp

  // minor3
  MOVUPS (R9), X3  // row0 -> minor3
  MULPS  X5, X3 // tmp * row0 (minor3)

  // tmp
  SHUFPS $0x4E, X5, X5

  // minor0
  MOVUPS 48(R9), X6 // row3 -> temporary
  MULPS  X5, X6 // tmp * row3 (temporary)
  SUBPS  X6, X0  // minor0 - temporary

  // minor3
  MOVUPS X3, X6       // minor3 -> temporary
  MOVUPS (R9), X3        // row0 -> minor3
  MULPS  X5, X3       // tmp * row0 (minor1)
  SUBPS  X6, X3       // minor1 - tmp
  SHUFPS $0x4E, X3, X3

  // ---------------------------------------

  // tmp
  MOVUPS 16(R9), X5         // row1 -> tmp
  SHUFPS $0x4E, X5, X5
  MULPS  48(R9), X5        // row3 * tmp
  SHUFPS $0xB1, X5, X5

  // row2
  MOVUPS 32(R9), X6
  SHUFPS $0x4E, X6, X6
  MOVUPS X6, 32(R9)

  // minor0
  MOVUPS 32(R9), X6 // row2 -> temporary
  MULPS  X5, X6 // tmp * row2 (temporary)
  ADDPS  X6, X0  // minor0 - temporary

  // minor2
  MOVUPS (R9), X2  // row2 -> minor2
  MULPS  X5, X2 // tmp * minor2

  // tmp
  SHUFPS $0x4E, X5, X5

  // minor0
  MOVUPS 32(R9), X6 // row2 -> temporary
  MULPS  X5, X6 // tmp * row2 (temporary)
  SUBPS  X6, X0  // minor0 - temporary

  // minor2
  MOVUPS X2, X6       // minor2 -> temporary
  MOVUPS (R9), X2        // row0 -> minor2
  MULPS  X5, X2       // tmp * row0 (minor1)
  SUBPS  X6, X2       // minor1 - tmp
  SHUFPS $0x4E, X2, X2

  // ---------------------------------------

  // tmp
  MOVUPS (R9), X5         // row0 -> tmp
  MULPS  16(R9), X5         // row1 * row2 (tmp)
  SHUFPS $0xB1, X5, X5

  // minor2
  MOVUPS X2, X6 // minor2 -> temporary
  MOVUPS 48(R9), X2 // row3 -> minor2
  MULPS  X5, X2 // tmp * row3 (minor2)
  ADDPS  X6, X2 // add

  // minor3
  MOVUPS X3, X6 // minor3 -> temporary
  MOVUPS 32(R9), X3 // row2 -> minor3
  MULPS  X5, X3 // tmp * row2 (minor3)
  SUBPS  X6, X3 // sub

  // tmp
  SHUFPS $0x4E, X5, X5

  // minor2
  MOVUPS X2, X6 // minor2 -> temporary
  MOVUPS 48(R9), X2 // row3 -> minor2
  MULPS  X5, X2 // tmp * row3 (minor2)
  SUBPS  X6, X2 // sub

  // minor3
  MOVUPS 32(R9), X6 // row2 -> temporary
  MULPS  X5, X6 // tmp * row2 (temporary)
  SUBPS  X6, X3  // sub

  // ---------------------------------------

  // tmp
  MOVUPS (R9), X5         // row0 -> tmp
  MULPS  48(R9), X5        // row3 * row0 (tmp)
  SHUFPS $0xB1, X5, X5

  // minor1
  MOVUPS 32(R9), X6 // row2 -> temporary
  MULPS  X5, X6 // tmp * row2 (temporary)
  SUBPS  X6, X1  // sub

  // minor2
  MOVUPS X2, X6 // minor2 -> temporary
  MOVUPS 16(R9), X2  // row1 -> minor2
  MULPS  X5, X2 // tmp * row3 (minor2)
  ADDPS  X6, X2 // add

  // tmp
  SHUFPS $0x4E, X5, X5

  // minor1
  MOVUPS X1, X6 // minor2 -> temporary
  MOVUPS 32(R9), X1 // row2 -> minor1
  MULPS  X5, X1 // tmp * row3 (minor1)
  ADDPS  X6, X1 // add

  // minor2
  MOVUPS 16(R9), X6  // row2 -> temporary
  MULPS  X5, X6 // tmp * row2 (temporary)
  SUBPS  X6, X2  // sub

  // ---------------------------------------

  // tmp
  MOVUPS (R9), X5         // row0 -> tmp
  MULPS  32(R9), X5        // row2 * row0 (tmp)
  SHUFPS $0xB1, X5, X5

  // minor1
  MOVUPS X1, X6 // minor1 -> temporary
  MOVUPS 48(R9), X1 // row3 -> minor1
  MULPS  X5, X1 // tmp * row3 (minor1)
  ADDPS  X6, X1 // sub

  // minor3
  MOVUPS 16(R9), X6  // row1 -> temporary
  MULPS  X5, X6 // tmp * row1 (temporary)
  SUBPS  X6, X3  // sub

  // tmp
  SHUFPS $0x4E, X5, X5

  // minor1
  MOVUPS 48(R9), X6 // row3 -> temporary
  MULPS  X5, X6 // tmp * row3 (temporary)
  SUBPS  X6, X1  // sub

  // minor3
  MOVUPS X3, X6 // minor3 -> temporary
  MOVUPS 16(R9), X3  // row1 -> minor3
  MULPS  X5, X3 // tmp * row1 (minor3)
  ADDPS  X6, X3 // add

  // ---------------------------------------

  // det
  MOVUPS (R9), X4         // row0 -> det
  MULPS  X0, X4         // minor0 * row0 (det)
  MOVUPS X4, X6        // det -> temporary
  SHUFPS $0x4E, X6, X6
  ADDPS  X6, X4
  MOVUPS X4, X6        // det -> temporary
  SHUFPS $0xB1, X6, X6
  ADDSS  X6, X4

  // break on det == 0
  MOVQ    $0, R9
  MOVQ    R9, X6
  UCOMISS X6, X4
  JNZ     result        // jump if not zero
  MOVB    $0, ret+8(FP) // return false
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
  MOVUPS X0, (R8) // write result

  // minor1
  MULPS  X4, X1
  MOVUPS X1, 16(R8) // write result

  // minor2
  MULPS  X4, X2
  MOVUPS X2, 32(R8) // write result

  // minor3
  MULPS  X4, X3
  MOVUPS X3, 48(R8) // write result

  // ---------------------------------------

  MOVB $1, R15
  MOVB R15, ret+8(FP) // return true
  RET
