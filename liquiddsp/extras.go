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
	"unsafe"

	"github.com/pkg/errors"
)

type Complexfloat C.liquid_float_complex

type FrameSyncCallback func(
	header []byte,
	headerValid bool,
	payload []byte,
	payloadValid bool,
	stats FrameSyncStats,
) bool

//export frameSyncCb
func frameSyncCb(
	header *C.uchar,
	headerValid C.int,
	payload *C.uchar,
	payloadLen C.uint,
	payloadValid C.int,
	stats C.framesyncstats_s,
	handle unsafe.Pointer,
) C.int {
	callback, ok := pointerHandles.Get(handle).(FrameSyncCallback)
	if !ok {
		panic("could not retrieve data for handle")
	}

	fmt.Println("                     CALLBACK payloadLen =", payloadLen)
	fmt.Println("                     CALLBACK payload =", payload)

	goHeader := C.GoBytes(unsafe.Pointer(header), 14)
	goHeaderValid := headerValid > 0
	goPayload := C.GoBytes(unsafe.Pointer(payload), C.int(payloadLen))
	goPayloadValid := payloadValid > 0
	var goStats FrameSyncStats
	goStats = *(*FrameSyncStats)(unsafe.Pointer(&stats))

	if callback(goHeader, goHeaderValid, goPayload, goPayloadValid, goStats) {
		return 0
	} else {
		return -1
	}
}

func maybeError(code C.int) error {
	if code != 0 {
		return errors.New(C.GoString(C.liquid_error_info(C.liquid_error_code(code))))
	}
	return nil
}
