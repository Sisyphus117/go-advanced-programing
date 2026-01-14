package qsort

/*
#include<stdlib.h>
typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
*/
import "C"
import "unsafe"

type SortFunc C.qsort_cmp_func_t

func Sort(base unsafe.Pointer, num, size int, cmp SortFunc) {
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(cmp))
}
