package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	dsp "github.com/redwood/liquiddsp-go"
)

func main() {
	// options
	var (
		k              uint32  = 4    // filter samples/symbol
		m              uint32  = 3    // filter delay (symbols)
		bt             float32 = 0.3  // bandwidth-time product
		numDataSymbols uint32  = 200  // number of data symbols
		SNRdB          float   = 30.0 // signal-to-noise ratio [dB]
		phi            float   = 0.0  // carrier phase offset
		dphi           float   = 0.0  // carrier frequency offset
	)

	// int dopt;
	// while ((dopt = getopt(argc,argv,"uhk:m:n:b:s:")) != EOF) {
	//     switch (dopt) {
	//     case 'u':
	//     case 'h': usage();              return 0;
	//     case 'k': k = atoi(optarg); break;
	//     case 'm': m = atoi(optarg); break;
	//     case 'n': numDataSymbols = atoi(optarg); break;
	//     case 'b': bt = atof(optarg); break;
	//     case 's': SNRdB = atof(optarg); break;
	//     default:
	//         exit(1);
	//     }
	// }

	// validate input
	if bt <= 0.0 || bt >= 1.0 {
		fmt.Fprintf(os.Stderr, "error: %s, bandwidth-time product must be in (0,1)\n", os.Args[0])
		os.Exit(1)
	}

	// derived values
	var (
		numSymbols uint32 = numDataSymbols + 2*m
		numSamples uint32 = k * numSymbols
		nstd              = math.Pow(10, -SNRdB/20) // noise standard deviation
	)

	// create modulator
	mod := dsp.GMSKModCreate(k, m, bt)
	dsp.GMSKModPrint(mod)

	// create demodulator
	demod := dsp.GMSKDemCreate(k, m, bt)
	dsp.GMSKDemSetEqBw(demod, 0.01)
	dsp.GMSKDemPrint(demod)

	var (
		i       uint32
		s       [numSymbols]uint32
		x       [numSamples]dsp.LiquidFloatComplex
		y       [numSamples]dsp.LiquidFloatComplex
		sym_out [numSymbols]uint32
	)

	// generate random data sequence
	for i := 0; i < numSymbols; i++ {
		s[i] = uint32(rand.Int31n(math.MaxUint32)) % 2
	}

	// modulate signal
	for i := 0; i < numSymbols; i++ {
		dsp.GMSKModModulate(mod, s[i], &x[k*i])
	}

	// add channel impairments
	// for i :=0; i<numSamples; i++ {
	//     y[i]  = x[i]*cexpf(_Complex_I*(phi + i*dphi))
	//     y[i] += nstd*(randnf() + _Complex_I*randnf())*M_SQRT1_2
	// }

	// demodulate signal
	for i = 0; i < numSymbols; i++ {
		dsp.GMSKDemDemodulate(demod, &y[k*i], &sym_out[i])
	}

	// destroy modem objects
	dsp.GMSKModDestroy(mod)
	dsp.GMSKDemDestroy(demod)

	// print results to screen
	delay := 2 * m
	numErrors := 0
	for i := delay; i < numSymbols; i++ {
		if s[i-delay] != sym_out[i] {
			numErrors += 1
		}
	}
	fmt.Printf("symbol errors : %4u / %4u\n", numErrors, numDataSymbols)

	//     // write results to output file
	//     FILE * fid = fopen(OUTPUT_FILENAME,"w");
	//     fprintf(fid,"%% %s : auto-generated file\n", OUTPUT_FILENAME);
	//     fprintf(fid,"clear all\n");
	//     fprintf(fid,"close all\n");
	//     fprintf(fid,"k = %u;\n", k);
	//     fprintf(fid,"m = %u;\n", m);
	//     fprintf(fid,"bt = %f;\n", bt);
	//     fprintf(fid,"numSymbols = %u;\n", numSymbols);
	//     fprintf(fid,"numSamples = %u;\n", numSamples);

	//     fprintf(fid,"x = zeros(1,numSamples);\n");
	//     fprintf(fid,"y = zeros(1,numSamples);\n");
	//     for (i=0; i<numSamples; i++) {
	//         fprintf(fid,"x(%4u) = %12.8f + j*%12.8f;\n", i+1, crealf(x[i]), cimagf(x[i]));
	//         fprintf(fid,"y(%4u) = %12.8f + j*%12.8f;\n", i+1, crealf(y[i]), cimagf(y[i]));
	//     }
	//     fprintf(fid,"t=[0:(numSamples-1)]/k;\n");
	//     fprintf(fid,"figure;\n");
	//     fprintf(fid,"plot(t,real(y),t,imag(y));\n");

	//     // artificially demodulate (generate receive filter, etc.)
	//     float hr[2*k*m+1];
	//     liquid_firdes_gmskrx(k,m,bt,0,hr);
	//     for (i=0; i<2*k*m+1; i++)
	//         fprintf(fid,"hr(%3u) = %12.8f;\n", i+1, hr[i]);
	//     fprintf(fid,"z = filter(hr,1,arg( ([y(2:end) 0]).*conj(y) ))/k;\n");
	//     fprintf(fid,"figure;\n");
	//     fprintf(fid,"plot(t,z,t(k:k:end),z(k:k:end),'or');\n");
	//     fprintf(fid,"grid on;\n");

	//     fclose(fid);
	//     printf("results written to '%s'\n", OUTPUT_FILENAME);

	//     return 0;
}
