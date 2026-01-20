#include "textflag.h"
// main 不需要接收参数，也不需要返回值，所以是 $16-0
// 16 字节是为了存放给 myprint 传递的 string (指针8 + 长度8)
TEXT ·main(SB), $16-0
	// 1. 从全局变量读取 string 的两个部分
	MOVQ ·helloworld+0(SB), AX // 获取数据指针
	MOVQ ·helloworld+8(SB), BX // 获取长度

	// 2. 将它们压入栈中，作为 myprint 的输入参数
	MOVQ AX, 0(SP) // 栈底存放指针
	MOVQ BX, 8(SP) // 紧接着存放长度

	// 3. 调用函数
	CALL ·myprint(SB)
	RET
