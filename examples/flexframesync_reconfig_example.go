package main

import (
	"fmt"
	"math/rand"
	"time"

	dsp "github.com/redwood/liquid-dsp/liquiddsp"
)

func main() {
	rand.Seed(time.Now().UTC().Unix())

	// define parameters
	// float SNRdB = 30.0f;
	// float noise_floor = -30.0f;
	const (
		numFrames  = 3
		SNRdB      = 30.0
		noiseFloor = -30.0
		bufLen     = 256
	)

	var (
		fgprops dsp.FlexFrameGenProps
		header  [14]byte
		// nstd    = float32(math.Pow(10.0, noiseFloor/20.0))
		// gamma   = float32(math.Pow(10.0, (SNRdB+noiseFloor)/20.0))
	)
	fgprops.InitDefault()
	fgprops.Check = uint32(dsp.LiquidCrcNone) // data validity check
	fgprops.Fec0 = uint32(dsp.LiquidFecNone)  // inner FEC scheme
	fgprops.Fec1 = uint32(dsp.LiquidFecNone)  // outer FEC scheme
	fgprops.ModScheme = uint32(dsp.LiquidModemQpsk)
	fg := dsp.NewFlexFrameGen(&fgprops)
	defer fg.Close()
	fs := dsp.NewFlexFrameSync(callback)
	defer fs.Close()

	// initialize header, payload
	for i := 0; i < 14; i++ {
		header[i] = byte(i)
	}

	// frame buffers, properties
	// var buf [bufLen]complex64

	payloads := []string{
		`{"foo": "bar", "baz": 123, "quux": false}`,
		`12329183u590ijslkdjf`,
		`alskdjf;laksjd;lfkjsal;dfkj`,
	}

	for j, payload := range payloads {
		// unsigned int noise_len = rand() % 100;
		// for (i=0; i<noise_len; i++) {
		//     buf[i] = myrandf(1.0f) + _Complex_I*myrandf(1.0f);
		// }
		// flexframesync_execute(fs, buf, buf_len);

		// configure frame generator properties
		// payloadLen := rand.Int()%256 + 1
		// dsp.FlexFrameGenPropsInitDefault(&fgprops)
		// fgprops.Check = uint32(dsp.LiquidCrcNone) // data validity check
		// fgprops.Fec0 = uint32(dsp.LiquidFecNone)  // inner FEC scheme
		// fgprops.Fec1 = uint32(dsp.LiquidFecNone)  // outer FEC scheme
		// // if rand.Int()%2 == 1 {
		// fgprops.ModScheme = uint32(dsp.LiquidModemQpsk)
		// } else {
		// 	fgprops.ModScheme = uint32(dsp.LiquidModemQam16)
		// }
		fg.SetProps(&fgprops)
		props, err := fg.Props()
		if err != nil {
			panic(err)
		}
		fmt.Printf("PROPS = %+v\n", props)

		// reallocate memory for payload
		// payload := payloadBackingBuffer[:payloadLen]

		// initialize payload
		// for i := 0; i < payloadLen; i++ {
		// 	payload[i] = byte(rand.Int() & 0xff)
		// }

		// set properties and assemble the frame

		// pl := make([]byte, len(payload))
		// for i := range pl {
		// 	pl[i] = byte(rand.Int() & 0xff)
		// }

		fg.Assemble(header[:], []byte(payload))
		fmt.Println("FRAME LEN =", fg.FrameLen())
		fmt.Printf("frame %d (payload len = %d), ", j, len(payload))
		fg.Print()

		if !fg.IsAssembled() {
			panic("not assembled")
		}

		// write the frame in blocks
		var frameComplete bool
		for !frameComplete {
			// write samples to buffer
			buf := make([]complex64, 256)
			frameComplete = fg.WriteSamples(buf)
			for i := range buf {
				fmt.Printf("%.15f + %.15fi\n", real(buf[i]), imag(buf[i]))
			}

			// add channel impairments (gain and noise)
			// for i := 0; i < buf_len; i++ {
			// 	buf[i] = buf[i]*gamma + nstd*(randnf()+_Complex_I*randnf())*M_SQRT1_2
			// }

			// push through sync
			fs.Execute(buf)
		}

		// noiseLen := rand.Int() % 100
		// for i:=0; i<noiseLen; i++ {
		//     buf[i] = myrandf(1.0f) + _Complex_I*myrandf(1.0f);
		// }
		// fs.Execute(buf, buf_len)

	} // num frames

	// print frame data statistics
	fs.Print()

	fmt.Printf("done.\n")
}

func callback(
	header []byte,
	headerValid bool,
	payload []byte,
	payloadValid bool,
	stats dsp.FrameSyncStats,
) bool {
	fmt.Println("******** callback invoked")

	// count bit errors (assuming all-zero message)
	var bitErrors uint32
	// for i := 0; i < len(payload); i++ {
	// 	bitErrors += dsp.LiquidCountOnes(uint32(payload[i]))
	// }

	stats.Print()
	fmt.Println("    payload             :   ", payload)
	fmt.Println("    header crc          :   ", headerValid)
	fmt.Println("    payload length      :   ", len(payload))
	fmt.Println("    payload crc         :   ", payloadValid)
	fmt.Println("    payload bit errors  :   ", bitErrors, "/", 8*len(payload))

	return false
}
