package main

import (
	"fmt"
	"math"
	"unsafe"

	dsp "github.com/redwood/liquid-dsp/liquiddsp"
)

func main() {
	msgs := []string{
		`{"foo": "bar", "baz": 123, "quux": false}`,
		`12329183u590ijslkdjf`,
		`alskdjf;laksjd;lfkjsal;dfkj`,
	}

	for _, msg := range msgs {
		var (
			ms    = dsp.LiquidModemQpsk    // mod. scheme
			check = dsp.LiquidCrc32        // data validity check
			fec0  = dsp.LiquidFecGolay2412 // fec (inner)
			fec1  = dsp.LiquidFecNone      // fec (outer)
			// SNRdB = float32(6.0)           // SNR [dB]
			// nstd  = float32(math.Pow(10.0, -float64(SNRdB)/20.0))

			payloadTx  = []byte(msg)
			payloadLen = uint32(len(payloadTx))
			payloadRx  = make([]byte, payloadLen)
		)

		// create and configure packet encoder/decoder object
		q := dsp.QpacketmodemCreate()
		dsp.QpacketmodemConfigure(q, payloadLen, check, fec0, fec1, int32(ms))
		dsp.QpacketmodemPrint(q)

		// get frame length
		frameLen := dsp.QpacketmodemGetFrameLen(q)
		fmt.Println("frame len", frameLen)

		// allocate memory for frame samples
		frameTx := make([]dsp.Complexfloat, frameLen)
		frameRx := make([]dsp.Complexfloat, frameLen)

		// encode frame
		if dsp.QpacketmodemEncode(q, payloadTx, frameTx) != 0 {
			panic("can't encode")
		}
		// for i := range frameTx {
		// 	frameTx[i] = frameTx[i] * 10000000000
		// }
		// fmt.Println(len(frameTx))
		// fmt.Println(frameTx[:100])

		// iqSamples := complex64sToIQSamples(frameTx)

		// add noise
		for i := uint32(0); i < frameLen; i++ {
			// x := complex(nstd, 0) * complex(dsp.Randnf(), dsp.Randnf()) * (1 / math.Sqrt2)
			// frameRx[i] = dsp.LiquidFloatComplex{Real: frameTx[i].Real + real(x), Imag: frameTx[i].Imag + imag(x)}
			// frameRx[i] = dsp.Complexfloat(complex64(frameTx[i]) + x)
			frameRx[i] = frameTx[i]
		}

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

	x := complex(float32(0.000000123), float32(0.000000456))
	y := dsp.Complexfloat(x)
	z := complex64(y)
	fmt.Println(z)
	fmt.Println(unsafe.Sizeof([1]dsp.Complexfloat{}))
	fmt.Println(unsafe.Sizeof([1]complex64{}))
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
		c64s[i] = complex(float32(i8s[i*2])/math.MaxInt8, float32(i8s[i*2+1])/math.MaxInt8)
	}
	return c64s
}
