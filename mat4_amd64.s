#include "textflag.h"

// func invertMat4SIMD(lhs *Mat4)
TEXT ·invertMat4SIMD(SB),NOSPLIT,$0
  // implementation based on:
  // http://www.intel.com/design/pentiumiii/sml/245043.htm

  /*
    X0 ... lhs row0
    X1 ... lhs row1
    X2 ... lhs row2
    x3 ... lhs row3

    X4 ... minor0
    X5 ... minor1
    X6 ... minor2
    X7 ... minor3

    X8 ... row0
    X9 ... row1
    X10 ... row2
    X11 ... row3

    X12 ... det
    X13 ... tmp
  */

  // load pointers into registers
  MOVQ lhs+0(FP), R8

  // load lhs
  /* MOVUPS (R8), X0   // row 1 */
  /* MOVUPS 16(R8), X1 // row 2 */
  /* MOVUPS 32(R8), X2 // row 3 */
  /* MOVUPS 48(R8), X3 // row 4 */

  // tmp1
  MOVUPS (R8), X13 // lhs0 -> tmp
  MOVUPS 16(R8), X1 // lhs1
  SHUFPS $0x44, X1, X13 // mask: 01 00 01 00

  // row1
  MOVUPS 32(R8), X9 // lhs2 -> row1
  MOVUPS 48(R8), X3 // lhs3
  SHUFPS $0x44, X3, X9 // mask: 01 00 01 00

  // row0
  MOVUPS X13, X8 // row1 -> row0
  SHUFPS $0x88, X9, X8 // mask: 10 00 10 00

  // row1
  SHUFPS $0xDD, X13, X9 // mask: 11 01 11 01

  // tmp1
  MOVUPS (R8), X13 // lhs0 -> tmp1
  // lhs1 still in X1
  SHUFPS $0xEE, X1, X13 // mask: 11 10 11 10

  // row3
  MOVUPS 32(R8), X11 // lhs2 -> row3
  // lhs3 still in X3
  SHUFPS $0xEE, X3, X11 // mask: 11 10 11 10

  // row2
  MOVUPS X13, X10 // row3 -> row2
  SHUFPS $0x88, X11, X10 // mask: 10 00 10 00

  // row3
  SHUFPS $0xDD, X13, X11

  // ---------------------------------------

  // tmp1
  MOVUPS X10, X13 // row2 -> tmp
  MULPS X11, X13 // row3 * row2 (tmp)
  SHUFPS $0xB1, X13, X13

  // minor0
  MOVUPS X9, X4 // row1 -> minor0
  MULPS X13, X4 // tmp * row1 (minor0)

  // minor1
  MOVUPS X8, X5 // row0 -> minor1
  MULPS X13, X5 // tmp * row0 (minor1)

  // tmp1
  SHUFPS $0x4E, X13, X13

  // minor0
  MOVUPS X4, X14 // temp minor0
  MOVUPS X9, X4 // row1 -> minor0
  MULPS X13, X4 // tmp * row1 (minor0)
  SUBPS X14, X4 // minor0 - tmp

  // minor1
  MOVUPS X5, X14 // temp minor1
  MOVUPS X8, X5 // row0 -> minor1
  MULPS X13, X5 // tmp * row0 (minor1)
  SUBPS X14, X5 // minor1 - tmp
  SHUFPS $0x4E, X5, X5

  // ---------------------------------------

  // tmp1
  MOVUPS X9, X13 // row1 -> tmp
  MULPS X10, X13 // row2 * row1 (tmp)
  SHUFPS $0xB1, X13, X13

  // minor0
  MOVUPS X4, X14 // temp minor0
  MOVUPS X11, X4 // row3 -> minor0
  MULPS X13, X4 // tmp * row3 (minor0)
  ADDPS X14, X4 // minor0 - tmp

  // minor3
  MOVUPS X8, X7 // row0 -> minor3
  MULPS X13, X7 // tmp * row0 (minor3)

  // tmp
  SHUFPS $0x4E, X13, X13

  // minor0
  MOVUPS X11, X14 // row3 -> temporary
  MULPS X13, X14 // tmp * row3 (temporary)
  SUBPS X14, X4 // minor0 - temporary

  // minor3
  MOVUPS X7, X14 // minor3 -> temporary
  MOVUPS X8, X7 // row0 -> minor3
  MULPS X13, X7 // tmp * row0 (minor1)
  SUBPS X14, X7 // minor1 - tmp
  SHUFPS $0x4E, X7, X7

  // ---------------------------------------

  // tmp
  MOVUPS X9, X13 // row1 -> tmp
  SHUFPS $0x4E, X13, X13
  MULPS X11, X13 // row3 * tmp
  SHUFPS $0xB1, X13, X13

  // row2
  SHUFPS $0x4E, X10, X10

  // minor0
  MOVUPS X10, X14 // row2 -> temporary
  MULPS X13, X14 // tmp * row2 (temporary)
  ADDPS X14, X4 // minor0 - temporary

  // minor2
  MOVUPS X8, X6 // row2 -> minor2
  MULPS X13, X6 // tmp * minor2

  // tmp
  SHUFPS $0x4E, X13, X13

  // minor0
  MOVUPS X10, X14 // row2 -> temporary
  MULPS X13, X14 // tmp * row2 (temporary)
  SUBPS X14, X4 // minor0 - temporary

  // minor2
  MOVUPS X6, X14 // minor2 -> temporary
  MOVUPS X8, X6 // row0 -> minor2
  MULPS X13, X6 // tmp * row0 (minor1)
  SUBPS X14, X6 // minor1 - tmp
  SHUFPS $0x4E, X6, X6

  // ---------------------------------------

  // tmp
  MOVUPS X8, X13 // row0 -> tmp
  MULPS X9, X13 // row1 * row2 (tmp)
  SHUFPS $0xB1, X13, X13

  // minor2
  MOVUPS X6, X14 // minor2 -> temporary
  MOVUPS X11, X6 // row3 -> minor2
  MULPS X13, X6 // tmp * row3 (minor2)
  ADDPS X14, X6 // add

  // minor3
  MOVUPS X7, X14 // minor3 -> temporary
  MOVUPS X10, X7 // row2 -> minor3
  MULPS X13, X7 // tmp * row2 (minor3)
  SUBPS X14, X7 // sub

  // tmp
  SHUFPS $0x4E, X13, X13

  // minor2
  MOVUPS X6, X14 // minor2 -> temporary
  MOVUPS X11, X6 // row3 -> minor2
  MULPS X13, X6 // tmp * row3 (minor2)
  SUBPS X14, X6 // sub

  // minor3
  MOVUPS X10, X14 // row2 -> temporary
  MULPS X13, X14 // tmp * row2 (temporary)
  SUBPS X14, X7 // sub

  // ---------------------------------------

  // tmp
  MOVUPS X8, X13 // row0 -> tmp
  MULPS X11, X13 // row3 * row0 (tmp)
  SHUFPS $0xB1, X13, X13

  // minor1
  MOVUPS X10, X14 // row2 -> temporary
  MULPS X13, X14 // tmp * row2 (temporary)
  SUBPS X14, X5 // sub

  // minor2
  MOVUPS X6, X14 // minor2 -> temporary
  MOVUPS X9, X6 // row1 -> minor2
  MULPS X13, X6 // tmp * row3 (minor2)
  ADDPS X14, X6 // add

  // tmp
  SHUFPS $0x4E, X13, X13

  // minor1
  MOVUPS X5, X14 // minor2 -> temporary
  MOVUPS X10, X5 // row2 -> minor1
  MULPS X13, X5 // tmp * row3 (minor1)
  ADDPS X14, X5 // add

  // minor2
  MOVUPS X9, X14 // row2 -> temporary
  MULPS X13, X14 // tmp * row2 (temporary)
  SUBPS X14, X6 // sub

  // ---------------------------------------

  // tmp
  MOVUPS X8, X13 // row0 -> tmp
  MULPS X10, X13 // row2 * row0 (tmp)
  SHUFPS $0xB1, X13, X13

  // minor1
  MOVUPS X5, X14 // minor1 -> temporary
  MOVUPS X11, X5 // row3 -> minor1
  MULPS X13, X5 // tmp * row3 (minor1)
  ADDPS X14, X5 // sub

  // minor3
  MOVUPS X9, X14 // row1 -> temporary
  MULPS X13, X14 // tmp * row1 (temporary)
  SUBPS X14, X7 // sub

  // tmp
  SHUFPS $0x4E, X13, X13

  // minor1
  MOVUPS X11, X14 // row3 -> temporary
  MULPS X13, X14 // tmp * row3 (temporary)
  SUBPS X14, X5 // sub

  // minor3
  MOVUPS X7, X14 // minor3 -> temporary
  MOVUPS X9, X7 // row1 -> minor3
  MULPS X13, X7 // tmp * row1 (minor3)
  ADDPS X14, X7 // add

  // ---------------------------------------

  // det
  MOVUPS X8, X12 // row0 -> det
  MULPS X4, X12 // minor0 * row0 (det)
  MOVUPS X12, X14 // det -> temporary
  SHUFPS $0x4E, X14, X14
  ADDPS X14, X12
  MOVUPS X12, X14 // det -> temporary
  SHUFPS $0xB1, X14, X14
  ADDSS X14, X12

  // break on det == 0
  MOVQ $0, R9
  MOVQ R9, X14
  UCOMISS X14, X12
  JNZ result // jump if not zero
  MOVB $0, ret+8(FP) // return false
  RET

result:

  // tmp
  MOVUPS X12, X13 // det -> tmp
  RCPSS X13, X13 // reciprical

  // det
  MOVUPS X13, X14 // tmp -> temporary
  MULSS X14, X12
  ADDSS X13, X13
  SUBSS X12, X13
  MOVUPS X13, X12
  SHUFPS $0x00, X12, X12

  // minor0
  MULPS X12, X4
  MOVUPS X4, (R8) // write result

  // minor1
  MULPS X12, X5
  MOVUPS X5, 16(R8) // write result

  // minor2
  MULPS X12, X6
  MOVUPS X6, 32(R8) // write result

  // minor3
  MULPS X12, X7
  MOVUPS X7, 48(R8) // write result

  // ---------------------------------------

  MOVB         $1, R15
  MOVB R15, ret+8(FP) // return true
  RET

// func mulMat4SIMD(lhs, rhs *Mat4)
TEXT ·mulMat4SIMD(SB),NOSPLIT,$0
  // load pointers into registers
  MOVQ lhs+0(FP), R8
  MOVQ rhs+8(FP), R9

  // load lhs into SSE registers
  MOVUPS (R8), X0   // row 1
  MOVUPS 16(R8), X1 // row 2
  MOVUPS 32(R8), X2 // row 3
  MOVUPS 48(R8), X3 // row 4

  ///  rhs row 1

  // mul lhs[0] with rhs[0][0]
  MOVUPS (R9), X4       // load rhs row 1
  SHUFPS $0x00, X4, X4  // broadcast 1. val
  MULPS X0, X4          // mul with lhs row 1

  // mul lhs[1] with rhs[0][1]
  MOVUPS (R9), X5       // load rhs row 1
  SHUFPS $0x55, X5, X5  // broadcast 2. val
  MULPS X1, X5          // mul with lhs row 2

  // mul lhs[2] with rhs[0][2]
  MOVUPS (R9), X6       // load rhs row 1
  SHUFPS $0xAA, X6, X6  // broadcast 3. val
  MULPS X2, X6          // mul with lhs row 3

  // mul lhs[3] with rhs[0][3]
  MOVUPS (R9), X7       // load rhs row 1
  SHUFPS $0xFF, X7, X7  // broadcast 4. val
  MULPS X3, X7          // mul with lhs row 4

  // add results
  ADDPS X5, X4
  ADDPS X6, X4
  ADDPS X7, X4

  // save result row
  MOVUPS X4, (R8)

  ///  rhs row 2

  // mul lhs[0] with rhs[1][0]
  MOVUPS 16(R9), X4       // load rhs row 2
  SHUFPS $0x00, X4, X4  // broadcast 1. val
  MULPS X0, X4          // mul with lhs row 1

  // mul lhs[1] with rhs[1][1]
  MOVUPS 16(R9), X5       // load rhs row 2
  SHUFPS $0x55, X5, X5  // broadcast 2. val
  MULPS X1, X5          // mul with lhs row 2

  // mul lhs[2] with rhs[1][2]
  MOVUPS 16(R9), X6       // load rhs row 2
  SHUFPS $0xAA, X6, X6  // broadcast 3. val
  MULPS X2, X6          // mul with lhs row 3

  // mul lhs[3] with rhs[1][3]
  MOVUPS 16(R9), X7       // load rhs row 2
  SHUFPS $0xFF, X7, X7  // broadcast 4. val
  MULPS X3, X7          // mul with lhs row 4

  // add results
  ADDPS X5, X4
  ADDPS X6, X4
  ADDPS X7, X4

  // save result row
  MOVUPS X4, 16(R8)

  ///  rhs row 3

  // mul lhs[0] with rhs[2][0]
  MOVUPS 32(R9), X4     // load rhs row 3
  SHUFPS $0x00, X4, X4  // broadcast 1. val
  MULPS X0, X4          // mul with lhs row 1

  // mul lhs[1] with rhs[2][1]
  MOVUPS 32(R9), X5     // load rhs row 3
  SHUFPS $0x55, X5, X5  // broadcast 2. val
  MULPS X1, X5          // mul with lhs row 2

  // mul lhs[2] with rhs[2][2]
  MOVUPS 32(R9), X6     // load rhs row 3
  SHUFPS $0xAA, X6, X6  // broadcast 3. val
  MULPS X2, X6          // mul with lhs row 3

  // mul lhs[3] with rhs[2][3]
  MOVUPS 32(R9), X7     // load rhs row 3
  SHUFPS $0xFF, X7, X7  // broadcast 4. val
  MULPS X3, X7          // mul with lhs row 4

  // add results
  ADDPS X5, X4
  ADDPS X6, X4
  ADDPS X7, X4

  // save result row
  MOVUPS X4, 32(R8)

  ///  rhs row 4

  // mul lhs[0] with rhs[3][0]
  MOVUPS 48(R9), X4     // load rhs row 3
  SHUFPS $0x00, X4, X4  // broadcast 1. val
  MULPS X0, X4          // mul with lhs row 1

  // mul lhs[1] with rhs[3][1]
  MOVUPS 48(R9), X5     // load rhs row 3
  SHUFPS $0x55, X5, X5  // broadcast 2. val
  MULPS X1, X5          // mul with lhs row 2

  // mul lhs[2] with rhs[3][2]
  MOVUPS 48(R9), X6     // load rhs row 3
  SHUFPS $0xAA, X6, X6  // broadcast 3. val
  MULPS X2, X6          // mul with lhs row 3

  // mul lhs[3] with rhs[3][3]
  MOVUPS 48(R9), X7     // load rhs row 3
  SHUFPS $0xFF, X7, X7  // broadcast 4. val
  MULPS X3, X7          // mul with lhs row 4

  // add results
  ADDPS X5, X4
  ADDPS X6, X4
  ADDPS X7, X4

  // save result row
  MOVUPS X4, 48(R8)

  /// done
  RET
