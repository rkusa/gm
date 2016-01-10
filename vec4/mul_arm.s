#include "textflag.h"

// func mulSIMD(lhs *Vec4, rhs float32)
TEXT ·mulSIMD(SB),NOSPLIT,$0
  B ·mul(SB)

