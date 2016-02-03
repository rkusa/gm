#include "textflag.h"

// func invertSIMD(lhs *Mat4)
TEXT ·invertSIMD(SB), NOSPLIT, $0
  // implementation based on:
  // http://www.intel.com/design/pentiumiii/sml/245043.htm

  // X0 ... minor0
  // X1 ... minor1
  // X2 ... minor2
  // X3 ... minor3
  //
  // X4 ... row0
  // X5 ... row1
  // X6 ... row2
  // X7 ... row3
  //
  // X8 ... det
  // X9 ... tmp
  // X10 ... tmp2

  // load pointers into registers
  //   (R8) ... lhs0
  // 16(R8) ... lhs1
  // 32(R8) ... lhs2
  // 48(R8) ... lhs3
  MOVQ lhs+0(FP), R8

  // tmp1
  MOVUPS (R8), X9          // lhs0 -> tmp
  SHUFPS $0x44, 16(R8), X9 // mask: 01 00 01 00

  // row1
  MOVUPS 32(R8), X5        // lhs2 -> row1
  SHUFPS $0x44, 48(R8), X5 // mask: 01 00 01 00

  // row0
  MOVUPS X9, X4       // row1 -> row0
  SHUFPS $0x88, X5, X4 // mask: 10 00 10 00

  // row1
  SHUFPS $0xDD, X9, X5 // mask: 11 01 11 01

  // tmp1
  MOVUPS (R8), X9          // lhs0 -> tmp1
  SHUFPS $0xEE, 16(R8), X9 // mask: 11 10 11 10

  // row3
  MOVUPS 32(R8), X7        // lhs2 -> row3
  SHUFPS $0xEE, 48(R8), X7 // mask: 11 10 11 10

  // row2
  MOVUPS X9, X6        // row3 -> row2
  SHUFPS $0x88, X7, X6 // mask: 10 00 10 00

  // row3
  SHUFPS $0xDD, X9, X7

  // ---------------------------------------

  // tmp1
  MOVUPS X6, X9        // row2 -> tmp
  MULPS  X7, X9        // row3 * row2 (tmp)
  SHUFPS $0xB1, X9, X9

  // minor0
  MOVUPS X5, X0  // row1 -> minor0
  MULPS  X9, X0 // tmp * row1 (minor0)

  // minor1
  MOVUPS X4, X1  // row0 -> minor1
  MULPS  X9, X1 // tmp * row0 (minor1)

  // tmp1
  SHUFPS $0x4E, X9, X9

  // minor0
  MOVUPS X0, X10 // temp minor0
  MOVUPS X5, X0  // row1 -> minor0
  MULPS  X9, X0 // tmp * row1 (minor0)
  SUBPS  X10, X0 // minor0 - tmp

  // minor1
  MOVUPS X1, X10       // temp minor1
  MOVUPS X4, X1        // row0 -> minor1
  MULPS  X9, X1       // tmp * row0 (minor1)
  SUBPS  X10, X1       // minor1 - tmp
  SHUFPS $0x4E, X1, X1

  // ---------------------------------------

  // tmp1
  MOVUPS X5, X9         // row1 -> tmp
  MULPS  X6, X9        // row2 * row1 (tmp)
  SHUFPS $0xB1, X9, X9

  // minor0
  MOVUPS X0, X10 // temp minor0
  MOVUPS X7, X0 // row3 -> minor0
  MULPS  X9, X0 // tmp * row3 (minor0)
  ADDPS  X10, X0 // minor0 - tmp

  // minor3
  MOVUPS X4, X3  // row0 -> minor3
  MULPS  X9, X3 // tmp * row0 (minor3)

  // tmp
  SHUFPS $0x4E, X9, X9

  // minor0
  MOVUPS X7, X10 // row3 -> temporary
  MULPS  X9, X10 // tmp * row3 (temporary)
  SUBPS  X10, X0  // minor0 - temporary

  // minor3
  MOVUPS X3, X10       // minor3 -> temporary
  MOVUPS X4, X3        // row0 -> minor3
  MULPS  X9, X3       // tmp * row0 (minor1)
  SUBPS  X10, X3       // minor1 - tmp
  SHUFPS $0x4E, X3, X3

  // ---------------------------------------

  // tmp
  MOVUPS X5, X9         // row1 -> tmp
  SHUFPS $0x4E, X9, X9
  MULPS  X7, X9        // row3 * tmp
  SHUFPS $0xB1, X9, X9

  // row2
  SHUFPS $0x4E, X6, X6

  // minor0
  MOVUPS X6, X10 // row2 -> temporary
  MULPS  X9, X10 // tmp * row2 (temporary)
  ADDPS  X10, X0  // minor0 - temporary

  // minor2
  MOVUPS X4, X2  // row2 -> minor2
  MULPS  X9, X2 // tmp * minor2

  // tmp
  SHUFPS $0x4E, X9, X9

  // minor0
  MOVUPS X6, X10 // row2 -> temporary
  MULPS  X9, X10 // tmp * row2 (temporary)
  SUBPS  X10, X0  // minor0 - temporary

  // minor2
  MOVUPS X2, X10       // minor2 -> temporary
  MOVUPS X4, X2        // row0 -> minor2
  MULPS  X9, X2       // tmp * row0 (minor1)
  SUBPS  X10, X2       // minor1 - tmp
  SHUFPS $0x4E, X2, X2

  // ---------------------------------------

  // tmp
  MOVUPS X4, X9         // row0 -> tmp
  MULPS  X5, X9         // row1 * row2 (tmp)
  SHUFPS $0xB1, X9, X9

  // minor2
  MOVUPS X2, X10 // minor2 -> temporary
  MOVUPS X7, X2 // row3 -> minor2
  MULPS  X9, X2 // tmp * row3 (minor2)
  ADDPS  X10, X2 // add

  // minor3
  MOVUPS X3, X10 // minor3 -> temporary
  MOVUPS X6, X3 // row2 -> minor3
  MULPS  X9, X3 // tmp * row2 (minor3)
  SUBPS  X10, X3 // sub

  // tmp
  SHUFPS $0x4E, X9, X9

  // minor2
  MOVUPS X2, X10 // minor2 -> temporary
  MOVUPS X7, X2 // row3 -> minor2
  MULPS  X9, X2 // tmp * row3 (minor2)
  SUBPS  X10, X2 // sub

  // minor3
  MOVUPS X6, X10 // row2 -> temporary
  MULPS  X9, X10 // tmp * row2 (temporary)
  SUBPS  X10, X3  // sub

  // ---------------------------------------

  // tmp
  MOVUPS X4, X9         // row0 -> tmp
  MULPS  X7, X9        // row3 * row0 (tmp)
  SHUFPS $0xB1, X9, X9

  // minor1
  MOVUPS X6, X10 // row2 -> temporary
  MULPS  X9, X10 // tmp * row2 (temporary)
  SUBPS  X10, X1  // sub

  // minor2
  MOVUPS X2, X10 // minor2 -> temporary
  MOVUPS X5, X2  // row1 -> minor2
  MULPS  X9, X2 // tmp * row3 (minor2)
  ADDPS  X10, X2 // add

  // tmp
  SHUFPS $0x4E, X9, X9

  // minor1
  MOVUPS X1, X10 // minor2 -> temporary
  MOVUPS X6, X1 // row2 -> minor1
  MULPS  X9, X1 // tmp * row3 (minor1)
  ADDPS  X10, X1 // add

  // minor2
  MOVUPS X5, X10  // row2 -> temporary
  MULPS  X9, X10 // tmp * row2 (temporary)
  SUBPS  X10, X2  // sub

  // ---------------------------------------

  // tmp
  MOVUPS X4, X9         // row0 -> tmp
  MULPS  X6, X9        // row2 * row0 (tmp)
  SHUFPS $0xB1, X9, X9

  // minor1
  MOVUPS X1, X10 // minor1 -> temporary
  MOVUPS X7, X1 // row3 -> minor1
  MULPS  X9, X1 // tmp * row3 (minor1)
  ADDPS  X10, X1 // sub

  // minor3
  MOVUPS X5, X10  // row1 -> temporary
  MULPS  X9, X10 // tmp * row1 (temporary)
  SUBPS  X10, X3  // sub

  // tmp
  SHUFPS $0x4E, X9, X9

  // minor1
  MOVUPS X7, X10 // row3 -> temporary
  MULPS  X9, X10 // tmp * row3 (temporary)
  SUBPS  X10, X1  // sub

  // minor3
  MOVUPS X3, X10 // minor3 -> temporary
  MOVUPS X5, X3  // row1 -> minor3
  MULPS  X9, X3 // tmp * row1 (minor3)
  ADDPS  X10, X3 // add

  // ---------------------------------------

  // det
  MOVUPS X4, X8         // row0 -> det
  MULPS  X0, X8         // minor0 * row0 (det)
  MOVUPS X8, X10        // det -> temporary
  SHUFPS $0x4E, X10, X10
  ADDPS  X10, X8
  MOVUPS X8, X10        // det -> temporary
  SHUFPS $0xB1, X10, X10
  ADDSS  X10, X8

  // break on det == 0
  MOVQ    $0, R9
  MOVQ    R9, X10
  UCOMISS X10, X8
  JNZ     result        // jump if not zero
  MOVB    $0, ret+8(FP) // return false
  RET

result:

  // tmp
  MOVUPS X8, X9 // det -> tmp
  RCPSS  X9, X9 // reciprical

  // det
  MOVUPS X9, X10        // tmp -> temporary
  MULSS  X10, X8
  ADDSS  X9, X9
  SUBSS  X8, X9
  MOVUPS X9, X8
  SHUFPS $0x00, X8, X8

  // minor0
  MULPS  X8, X0
  MOVUPS X0, (R8) // write result

  // minor1
  MULPS  X8, X1
  MOVUPS X1, 16(R8) // write result

  // minor2
  MULPS  X8, X2
  MOVUPS X2, 32(R8) // write result

  // minor3
  MULPS  X8, X3
  MOVUPS X3, 48(R8) // write result

  // ---------------------------------------

  MOVB $1, R15
  MOVB R15, ret+8(FP) // return true
  RET
