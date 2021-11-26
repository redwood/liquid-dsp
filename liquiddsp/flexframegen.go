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
#include "extras.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type FlexFrameGen struct {
	ptr C.flexframegen
}

// FlexFrameGenCreate function as declared in include/liquid.h:5120
func NewFlexFrameGen(props *FlexFrameGenProps) *FlexFrameGen {
	ptr := C.flexframegen_create((*C.flexframegenprops_s)(unsafe.Pointer(props)))
	runtime.KeepAlive(props)
	return &FlexFrameGen{ptr: ptr}
}

// FlexFrameGenDestroy function as declared in include/liquid.h:5123
func (fg *FlexFrameGen) Close() error {
	err := C.flexframegen_destroy(fg.ptr)
	runtime.KeepAlive(fg)
	return maybeError(err)
}

// FlexFrameGenGetProps function as declared in include/liquid.h:5135
func (fg *FlexFrameGen) Props() (*FlexFrameGenProps, error) {
	var props FlexFrameGenProps
	err := C.flexframegen_getprops(fg.ptr, (*C.flexframegenprops_s)(unsafe.Pointer(&props)))
	runtime.KeepAlive(fg)
	return &props, maybeError(err)
}

// FlexFrameGenSetProps function as declared in include/liquid.h:5138
func (fg *FlexFrameGen) SetProps(props *FlexFrameGenProps) error {
	err := C.flexframegen_setprops(fg.ptr, (*C.flexframegenprops_s)(unsafe.Pointer(props)))
	runtime.KeepAlive(fg)
	return maybeError(err)
}

// FlexFrameGenSetHeaderLen function as declared in include/liquid.h:5141
func (fg *FlexFrameGen) SetHeaderLen(length uint32) error {
	err := C.flexframegen_set_header_len(fg.ptr, C.uint(length))
	runtime.KeepAlive(fg)
	return maybeError(err)
}

// FlexFrameGenSetHeaderProps function as declared in include/liquid.h:5144
func (fg *FlexFrameGen) SetHeaderProps(props *FlexFrameGenProps) error {
	err := C.flexframegen_set_header_props(fg.ptr, (*C.flexframegenprops_s)(unsafe.Pointer(props)))
	runtime.KeepAlive(fg)
	return maybeError(err)
}

// FlexFrameGenGetFrameLen function as declared in include/liquid.h:5148
func (fg *FlexFrameGen) FrameLen() uint32 {
	x := uint32(C.flexframegen_getframelen(fg.ptr))
	runtime.KeepAlive(fg)
	return x
}

// FlexFrameGenWriteSamples function as declared in include/liquid.h:5166
func (fg *FlexFrameGen) WriteSamples(buffer []complex64) (frameComplete bool) {
	buf := make([]C.liquid_float_complex, len(buffer))
	fc := C.flexframegen_write_samples(
		fg.ptr,
		(*C.liquid_float_complex)(&(buf[0])),
		C.uint(len(buffer)),
	)
	fmt.Println("LEN =", len(buf))
	for i := range buf {
		buffer[i] = complex(float32(C.crealf(buf[i])), float32(C.cimagf(buf[i])))
	}
	runtime.KeepAlive(fg)
	runtime.KeepAlive(buf)
	runtime.KeepAlive(buffer)
	return fc > 0
}

// FlexFrameGenAssemble function as declared in include/liquid.h:5155
func (fg *FlexFrameGen) Assemble(header []byte, payload []byte) error {
	err := C.flexframegen_assemble(
		fg.ptr,
		(*C.uchar)(&(header[0])),
		(*C.uchar)(&(payload[0])),
		C.uint(len(payload)),
	)
	runtime.KeepAlive(fg)
	return maybeError(err)
}

// FlexFrameGenPrint function as declared in include/liquid.h:5126
func (fg *FlexFrameGen) Print() error {
	err := C.flexframegen_print(fg.ptr)
	runtime.KeepAlive(fg)
	return maybeError(err)
}

// FlexFrameGenReset function as declared in include/liquid.h:5129
func (fg *FlexFrameGen) Reset() error {
	err := C.flexframegen_reset(fg.ptr)
	runtime.KeepAlive(fg)
	return maybeError(err)
}

// FlexFrameGenIsAssembled function as declared in include/liquid.h:5132
func (fg *FlexFrameGen) IsAssembled() bool {
	is := C.flexframegen_is_assembled(fg.ptr)
	runtime.KeepAlive(fg)
	return is > 0
}
