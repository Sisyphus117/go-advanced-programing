package gls

import (
	"sync"
	"unsafe"
)

var gls struct {
	m map[int64]map[any]any
	sync.Mutex
}

func init() {
	gls.m = make(map[int64]map[any]any)
}

const g_goid_offset = 152

func getg() unsafe.Pointer

func GetGoroutineId() int64 {
	g := getg()
	p := (*int64)(unsafe.Add(g, g_goid_offset))
	return *p
}

func getMap() map[any]any {
	gls.Lock()
	defer gls.Unlock()

	goid := GetGoroutineId()
	if m, _ := gls.m[goid]; m != nil {
		return m
	}
	m := make(map[any]any)
	gls.m[goid] = m
	return m
}

func Get(key any) any {
	return getMap()[key]
}

func Put(key any, value any) {
	getMap()[key] = value
}

func Delete(key any) {
	delete(getMap(), key)
}

func Clean() {
	gls.Lock()
	defer gls.Unlock()

	delete(gls.m, GetGoroutineId())
}
