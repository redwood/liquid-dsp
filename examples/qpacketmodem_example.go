package main

import (
	"fmt"
	"math"

	dsp "github.com/redwood/liquid-dsp/liquiddsp"
)

func main() {
	var (
		ms    = dsp.LiquidModemQpsk    // mod. scheme
		check = dsp.LiquidCrc32        // data validity check
		fec0  = dsp.LiquidFecGolay2412 // fec (inner)
		fec1  = dsp.LiquidFecNone      // fec (outer)
		SNRdB = float32(6.0)           // SNR [dB]
		nstd  = float32(math.Pow(10.0, -float64(SNRdB)/20.0))

		payloadTx  = []byte(`{"foo": "bar", "baz": 123, "quux": false}`)
		payloadLen = uint32(len(payloadTx))
		payloadRx  = make([]byte, payloadLen)
	)

	// create and configure packet encoder/decoder object
	q := dsp.QpacketmodemCreate()
	dsp.QpacketmodemConfigure(q, payloadLen, check, fec0, fec1, int32(ms))
	dsp.QpacketmodemPrint(q)

	// get frame length
	frameLen := dsp.QpacketmodemGetFrameLen(q)

	// allocate memory for frame samples
	frameTx := make([]complex64, frameLen)
	frameRx := make([]complex64, frameLen)

	// encode frame
	dsp.QpacketmodemEncode(q, string(payloadTx), frameTx)

	iqSamples := complex64sToIQSamples(frameTx)

	// // add noise
	// for i := uint32(0); i < frameLen; i++ {
	// 	frameRx[i] = frameTx[i] + complex(nstd, 0)*complex(dsp.Randnf(), dsp.Randnf())*(1/math.Sqrt2)
	// }

	// decode frame
	crcPass := dsp.QpacketmodemDecode(q, frameRx, payloadRx) > 0

	// count errors
	numBitErrors := dsp.CountBitErrorsArray(payloadTx, payloadRx, payloadLen)

	// print results
	if crcPass {
		fmt.Printf("payload PASS, errors: %d / %d\n", numBitErrors, 8*payloadLen)
	} else {
		fmt.Printf("payload FAIL, errors: %d / %d\n", numBitErrors, 8*payloadLen)
	}

	// destroy allocated objects
	dsp.QpacketmodemDestroy(q)

	fmt.Println("output =", string(payloadRx))

}

func complex64sToIQSamples(c64s []complex64) []int8 {
	i8s := make([]int8, len(c64s)*2)
	for i := 0; i < len(c64s); i++ {
		i8s[i*2] = int8(imag(c64s[i]) * math.MaxInt8)
		i8s[i*2+1] = int8(real(c64s[i]) * math.MaxInt8)
	}
	return i8s
}

func iqSamplesToComplexFloat32s(i8s []int8) []complex64 {
	c64s := make([]complex64, len(i8s)/2)
	for i := 0; i < len(i8s)/2; i++ {
		c64s[i] = complex(i8s[i*2]/math.MaxInt8, i8s[i*2+1]/math.MaxInt8)
	}
	return c64s
}
