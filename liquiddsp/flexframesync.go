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
	"runtime"
	"unsafe"
)

type FlexFrameSync struct {
	ptr C.flexframesync
}

func NewFlexFrameSync(callback FrameSyncCallback) *FlexFrameSync {
	handle := pointerHandles.Track(callback)
	ccallback := (*[0]byte)(C.frame_sync_callback_helper)
	ptr := C.flexframesync_create(ccallback, handle)
	return &FlexFrameSync{ptr: ptr}
}

// FlexFrameSyncDestroy function as declared in include/liquid.h:5181
func (fs *FlexFrameSync) Close() error {
	err := C.flexframesync_destroy(fs.ptr)
	runtime.KeepAlive(fs)
	return maybeError(err)
}

// FlexFrameSyncExecute function as declared in include/liquid.h:5212
func (fs *FlexFrameSync) Execute(x []complex64) error {
	buf := make([]C.liquid_float_complex, len(x))
	for i := range buf {
		buf[i] = C.make_complex_float(C.float(real(x[i])), C.float(imag(x[i]))) //complex(float32(C.crealf(x[i])), float32(C.cimagf(x[i])))
	}
	err := C.flexframesync_execute(
		fs.ptr,
		(*C.liquid_float_complex)(&(buf[0])),
		C.uint(len(x)),
	)
	return maybeError(err)
}

// FlexFrameSyncReset function as declared in include/liquid.h:5187
func (fs *FlexFrameSync) Reset() error {
	err := C.flexframesync_reset(fs.ptr)
	runtime.KeepAlive(fs)
	return maybeError(err)
}

// FlexFrameSyncIsFrameOpen function as declared in include/liquid.h:5190
func (fs *FlexFrameSync) IsFrameOpen() bool {
	is := C.flexframesync_is_frame_open(fs.ptr)
	runtime.KeepAlive(fs)
	return is > 0
}

// FlexFrameSyncSetHeaderLen function as declared in include/liquid.h:5193
func (fs *FlexFrameSync) SetHeaderLen(length uint32) error {
	err := C.flexframesync_set_header_len(fs.ptr, C.uint(length))
	runtime.KeepAlive(fs)
	return maybeError(err)
}

// FlexFrameSyncDecodeHeaderSoft function as declared in include/liquid.h:5197
func (fs *FlexFrameSync) DecodeHeaderSoft(soft int32) error {
	err := C.flexframesync_decode_header_soft(fs.ptr, C.int(soft))
	runtime.KeepAlive(fs)
	return maybeError(err)
}

// FlexFrameSyncDecodePayloadSoft function as declared in include/liquid.h:5201
func (fs *FlexFrameSync) DecodePayloadSoft(soft int32) error {
	err := C.flexframesync_decode_payload_soft(fs.ptr, C.int(soft))
	runtime.KeepAlive(fs)
	return maybeError(err)
}

// FlexFrameSyncSetHeaderProps function as declared in include/liquid.h:5205
func (fs *FlexFrameSync) SetHeaderProps(props *FlexFrameGenProps) error {
	err := C.flexframesync_set_header_props(fs.ptr, (*C.flexframegenprops_s)(unsafe.Pointer(props)))
	runtime.KeepAlive(fs)
	runtime.KeepAlive(props)
	return maybeError(err)
}

// FlexFrameSyncResetFrameDataStats function as declared in include/liquid.h:5217
func (fs *FlexFrameSync) ResetFrameDataStats() error {
	err := C.flexframesync_reset_framedatastats(fs.ptr)
	runtime.KeepAlive(fs)
	return maybeError(err)
}

// FlexFrameSyncGetFrameDataStats function as declared in include/liquid.h:5218
// func (fs *FlexFrameSync) FrameDataStats() FrameDataStats {
// 	stats := C.flexframesync_get_framedatastats(fs.ptr)
//     runtime.KeepAlive(fs)
// 	return
// }

// FlexFrameSyncPrint function as declared in include/liquid.h:5184
func (fs *FlexFrameSync) Print() error {
	err := C.flexframesync_print(fs.ptr)
	runtime.KeepAlive(fs)
	return maybeError(err)
}

// FlexFrameSyncDebugEnable function as declared in include/liquid.h:5221
func (fs *FlexFrameSync) DebugEnable() error {
	err := C.flexframesync_debug_enable(fs.ptr)
	runtime.KeepAlive(fs)
	return maybeError(err)
}

// FlexFrameSyncDebugDisable function as declared in include/liquid.h:5222
func (fs *FlexFrameSync) DebugDisable() error {
	err := C.flexframesync_debug_disable(fs.ptr)
	runtime.KeepAlive(fs)
	return maybeError(err)
}

// FlexFrameSyncDebugPrint function as declared in include/liquid.h:5223
// func (fs *FlexFrameSync) DebugPrint(filename string) error {
// 	cfilename, cfilenameAllocMap := unpackPCharString(filename)
// 	err := C.flexframesync_debug_print(fs.ptr, cfilename)
// 	runtime.KeepAlive(fs)
// 	return maybeError(err)
// }
