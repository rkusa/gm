#include "textflag.h"

// func divSIMD(lhs *Vec4, rhs float32)
TEXT ·divSIMD(SB),NOSPLIT,$0
  B ·div(SB)

