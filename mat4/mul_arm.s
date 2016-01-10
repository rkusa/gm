#include "textflag.h"

// func mulSIMD(out, lhs, rhs *Mat4)
TEXT ·mulSIMD(SB),NOSPLIT,$0
  B ·mul(SB)

