#include "textflag.h"
GLOBL text<>(SB), NOPTR, $16
DATA text<>+0(SB)/8, $"Hello Wo"
DATA text<>+8(SB)/8, $"rld!"

GLOBL ·helloworld(SB), NOPTR, $16
DATA ·helloworld+0(SB)/8, $text<>(SB)
DATA ·helloworld+8(SB)/8, $12
