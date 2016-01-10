#include "textflag.h"

// func subSIMD(lhs, rhs *Vec4)
TEXT ·subSIMD(SB),NOSPLIT,$0
  B ·sub(SB)

