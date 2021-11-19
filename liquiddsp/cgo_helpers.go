// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 18 Nov 2021 16:40:08 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package liquiddsp

/*
#cgo CFLAGS: -msse4.1 -Wall -fPIC -I. -I.. -I../include
#cgo LDFLAGS: -lm ${SRCDIR}/../libliquid.a
#include <inttypes.h>
#include <complex.h>
#include "liquid.h"
#include "liquid.internal.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// unpackPUcharString represents the data from Go string as *C.uchar and avoids copying.
func unpackPUcharString(str string) (*C.uchar, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.uchar)(h.Data), cgoAllocsUnknown
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// allocComplex64Memory allocates memory for type C.complexfloat in C.
// The caller is responsible for freeing the this memory via C.free.
func allocComplex64Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfComplex64Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfComplex64Value = unsafe.Sizeof([1]C.complexfloat{})

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackArgSComplex64 transforms a sliced Go data structure into plain C format.
func unpackArgSComplex64(x []complex64) (unpacked *C.complexfloat, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.complexfloat) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocComplex64Memory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: mem0,
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.complexfloat)(unsafe.Pointer(h0))
	for i0 := range x {
		v0[i0] = C.complexfloat(x[i0])
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.complexfloat)(h.Data)
	return
}

// packSComplex64 reads sliced Go data structure out from plain C format.
func packSComplex64(v []complex64, ptr0 *C.complexfloat) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfComplex64Value]C.complexfloat)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = complex64(ptr1)
	}
}
