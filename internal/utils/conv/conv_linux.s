#include "textflag.h"
#include "go_asm.h"

TEXT ·fastb(SB), NOSPLIT|NOFRAME|DUPOK, $0
   MOVQ 0(SP), DX
   MOVQ 8(SP), AX
   RET

