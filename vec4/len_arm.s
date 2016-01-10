#include "textflag.h"

// func lenSIMD(lhs *Vec4) float32
TEXT ·lenSIMD(SB),NOSPLIT,$0
  B ·len(SB)

