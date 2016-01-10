#include "textflag.h"

// func addSIMD(lhs, rhs *Vec4)
TEXT ·addSIMD(SB),NOSPLIT,$0
  B ·add(SB)

