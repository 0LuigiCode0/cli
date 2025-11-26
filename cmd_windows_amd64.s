#include "textflag.h"

TEXT ·write(SB), NOSPLIT, $0-32
   MOVQ fd+0(FP), R10
   MOVQ $0, DX
   MOVQ $0, R8
   MOVQ $0, R9
   MOVQ ioStatus+8(FP), AX
   MOVQ AX, 40(SP) 
   MOVQ buff+16(FP), AX
   MOVQ AX, 48(SP)
   MOVQ len+24(FP), AX
   MOVQ AX, 56(SP)
   MOVQ $0, 64(SP)
   MOVQ $0, 72(SP)
   
   MOVQ $0x0008, AX
   SYSCALL
   RET
   