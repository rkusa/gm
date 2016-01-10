#include "textflag.h"

// func invertSIMD(lhs *Mat4)
TEXT ·invertSIMD(SB),NOSPLIT,$0
  B ·invert(SB)

