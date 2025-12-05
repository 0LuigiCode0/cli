#include "textflag.h"

#define IsC $0x01
#define F1 $0x02
#define F2 $0x04
#define F3 $0x08
#define F4 $0x10
#define F5 $0x20
#define F6 $0x40
#define F7 $0x80
#define F8 $0x100
#define F9 $0x200
#define F10 $0x400
#define F11 $0x800
#define F12 $0x1000
#define F13 $0x2000
#define F14 $0x4000
#define F15 $0x8000
#define FOut $0x10000

// for test
TEXT ·call3(SB),NOSPLIT, $0-48
   MOVQ mask+8(FP), R11
   JMP v1
   // TESTQ $0x01, R11
   // JNZ c

   // MOVQ func+0(FP), AX
   // MOVQ x+16(FP), DI
   // MOVQ x2+24(FP), SI
   // MOVQ x3+32(FP), DX
   // SYSCALL
   // JMP retX

   c:
   //целочисленные
   x:
   x1:
      JMP AX
   x2:
      JMP AX
   x3:
      JMP AX
   //с точкой
   f:
   f1:
      JMP AX
   f2:
      JMP AX
   f3:
      JMP AX
   //переменные
      v1:
      LEAQ call(SB), AX
      JMP x1
   //вызов
   call: 
      MOVQ func+0(FP), R10
   CALL R10

   retX: TESTQ FOut, mask+8(FP)
         JNZ retF
         MOVQ AX, r+40(FP)
         RET
   retF: MOVQ X0, r+40(FP)
         RET

// call:
//    MOVQ func+0(FP), R10
//    CALL R10
//    RET

// #define init\
//    MOVQ mask+8(FP), R11\
//    TESTQ $0x01, R11\
//    JNZ c

// #define syscall\
//    MOVQ func+0(FP), AX\
//    SYSCALL\
//    JMP retX

// #define checkFloat(mask,lbl,shift,reg)\
//    TESTQ mask, R11\
//    JZ lbl\
//    MOVQ x+shift(FP), reg

// #define checkCall\ 
//    MOVQ func+0(FP), R10\
//    TESTQ $0x0F, SP\
//    JNZ grow\
//    CALL R10\
//    JMP retX

// #define growCall\
//    SUBQ $8, SP\
//    CALL R10\
//    ADDQ $8, SP

// #define checkRet(shift)\
//    TESTQ $0x100, mask+8(FP)\
//    JNZ retF\
//    ret(AX,shift)

// #define ret(reg,shift)\
//    MOVQ reg, r+shift(FP)\
//    RET

// TEXT ·call3(SB), NOSPLIT, $0-48
//    MOVQ x1+16(FP), DI
//    MOVQ x2+24(FP), SI
//    MOVQ x3+32(FP), DX

//    init;syscall
//    c:
//       x1:   checkFloat($0x02,x2,16,X0)
//       x2:   checkFloat($0x04,x3,24,X1)
//       x3:   checkFloat($0x08,call,32,X2)
//       call: checkCall
//       grow: growCall
//    retX: checkRet(40)
//    retF: ret(X0,40)

// TEXT ·call6(SB), NOSPLIT, $0-72
//    MOVQ x1+16(FP), DI
//    MOVQ x2+24(FP), SI
//    MOVQ x3+32(FP), DX
//    MOVQ x5+48(FP), R8
//    MOVQ x6+56(FP), R9

//    init
//    MOVQ x4+40(FP), R10;syscall
//    c: MOVQ x4+40(FP), CX
//       x1:   checkFloat($0x02,x2,16,X0)
//       x2:   checkFloat($0x04,x3,24,X1)
//       x3:   checkFloat($0x08,x4,32,X2)
//       x4:   checkFloat($0x10,x5,40,X3)
//       x5:   checkFloat($0x20,x6,48,X4)
//       x6:   checkFloat($0x40,call,56,X5)
//       call: checkCall
//       grow: growCall
//    retX: checkRet(64)
//    retF: ret(X0,64)

// TEXT ·call9(SB), NOSPLIT, $72-96
//    MOVQ x1+16(FP), DI
//    MOVQ x2+24(FP), SI
//    MOVQ x3+32(FP), DX
//    MOVQ x5+48(FP), R8
//    MOVQ x6+56(FP), R9
//    MOVQ x7+64(FP), AX; MOVQ AX, 40(SP)
//    MOVQ x8+72(FP), AX; MOVQ AX, 48(SP)
//    MOVQ x9+80(FP), AX; MOVQ AX, 56(SP)

//    init;
//    MOVQ x4+40(FP), R10;syscall
//    c: MOVQ x4+40(FP), CX
//       x1:   checkFloat($0x02,x2,16,X0)
//       x2:   checkFloat($0x04,x3,24,X1)
//       x3:   checkFloat($0x08,x4,32,X2)
//       x4:   checkFloat($0x10,x5,40,X3)
//       x5:   checkFloat($0x20,x6,48,X4)
//       x6:   checkFloat($0x40,call,56,X5)
//       call: checkCall
//       grow: growCall
//    retX: checkRet(64)
//    retF: ret(X0,64)
//    MOVQ func+0(FP), AX
//    CMPQ isc+8(FP), $1
//    JE c

//    MOVQ x4+40(FP), R10
//    SYSCALL
//    JNE ret
// c:
//    MOVQ x4+40(FP), CX
//    CALL AX
// ret:
//    MOVQ AX, ret+88(FP)
   // RET

// TEXT ·call12(SB), NOSPLIT, $96-120
//    MOVQ x1+16(FP), CX
//    MOVQ x2+24(FP), DX
//    MOVQ x3+32(FP), R8
//    MOVQ x5+48(FP), AX; MOVQ AX, 32(SP)
//    MOVQ x6+56(FP), AX; MOVQ AX, 40(SP)
//    MOVQ x7+64(FP), AX; MOVQ AX, 48(SP)
//    MOVQ x8+72(FP), AX; MOVQ AX, 56(SP)
//    MOVQ x9+80(FP), AX; MOVQ AX, 64(SP)
//    MOVQ x10+88(FP), AX; MOVQ AX, 72(SP)
//    MOVQ x11+96(FP), AX; MOVQ AX, 80(SP)
//    MOVQ x12+104(FP), AX; MOVQ AX, 88(SP)

//    MOVQ func+0(FP), AX
//    CMPQ isc+8(FP), $1
//    JE c

//    MOVQ x4+40(FP), R10
//    SYSCALL
//    JNE ret
// c:
//    MOVQ x4+40(FP), CX
//    CALL AX
// ret:
//    MOVQ AX, ret+112(FP)
//    RET

// TEXT ·call15(SB), NOSPLIT, $120-144
//    MOVQ x1+16(FP), CX
//    MOVQ x2+24(FP), DX
//    MOVQ x3+32(FP), R8
//    MOVQ x5+48(FP), AX; MOVQ AX, 32(SP)
//    MOVQ x6+56(FP), AX; MOVQ AX, 40(SP)
//    MOVQ x7+64(FP), AX; MOVQ AX, 48(SP)
//    MOVQ x8+72(FP), AX; MOVQ AX, 56(SP)
//    MOVQ x9+80(FP), AX; MOVQ AX, 64(SP)
//    MOVQ x10+88(FP), AX; MOVQ AX, 72(SP)
//    MOVQ x11+96(FP), AX; MOVQ AX, 80(SP)
//    MOVQ x12+104(FP), AX; MOVQ AX, 88(SP)
//    MOVQ x13+112(FP), AX; MOVQ AX, 96(SP)
//    MOVQ x14+120(FP), AX; MOVQ AX, 104(SP)
//    MOVQ x15+128(FP), AX; MOVQ AX, 112(SP)

//    MOVQ func+0(FP), AX
//    CMPQ isc+8(FP), $1
//    JE c

//    MOVQ x4+40(FP), R10
//    SYSCALL
//    JNE ret
// c:
//    MOVQ x4+40(FP), CX
//    CALL AX
// ret:
//    MOVQ AX, ret+136(FP)
//    RET

// TEXT ·call18(SB), NOSPLIT, $144-168
//    MOVQ x1+16(FP), CX
//    MOVQ x2+24(FP), DX
//    MOVQ x3+32(FP), R8
//    MOVQ x5+48(FP), AX; MOVQ AX, 32(SP)
//    MOVQ x6+56(FP), AX; MOVQ AX, 40(SP)
//    MOVQ x7+64(FP), AX; MOVQ AX, 48(SP)
//    MOVQ x8+72(FP), AX; MOVQ AX, 56(SP)
//    MOVQ x9+80(FP), AX; MOVQ AX, 64(SP)
//    MOVQ x10+88(FP), AX; MOVQ AX, 72(SP)
//    MOVQ x11+96(FP), AX; MOVQ AX, 80(SP)
//    MOVQ x12+104(FP), AX; MOVQ AX, 88(SP)
//    MOVQ x13+112(FP), AX; MOVQ AX, 96(SP)
//    MOVQ x14+120(FP), AX; MOVQ AX, 104(SP)
//    MOVQ x15+128(FP), AX; MOVQ AX, 112(SP)
//    MOVQ x16+136(FP), AX; MOVQ AX, 120(SP)
//    MOVQ x17+144(FP), AX; MOVQ AX, 128(SP)
//    MOVQ x18+152(FP), AX; MOVQ AX, 136(SP)

//    MOVQ func+0(FP), AX
//    CMPQ isc+8(FP), $1
//    JE c

//    MOVQ x4+40(FP), R10
//    SYSCALL
//    JNE ret
// c:
//    MOVQ x4+40(FP), CX
//    CALL AX
// ret:
//    MOVQ AX, ret+160(FP)
   // RET



// Архитектура	Volatile регистры	Назначение
// x64 Windows	RAX, RCX, RDX, R8, R9, R10, R11, XMM0-XMM5	Можно использовать без сохранения
// x64 Linux	RAX, RDI, RSI, RDX, RCX, R8, R9, R10, R11, XMM0-XMM15	Можно использовать свободно
// 2. Non-volatile (Callee-saved) регистры - нужно сохранять
// Эти регистры должны быть сохранены если вы их изменяете:

// Архитектура	Non-volatile регистры
// x64 Windows	RBX, RBP, RDI, RSI, RSP, R12-R15, XMM6-XMM15
// x64 Linux	RBX, RBP, R12-R15