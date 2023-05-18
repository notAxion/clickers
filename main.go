package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

const (
	Duration   = 2
	SampleRate = 44100
	Frequency  = 440
)

func main() {
	outputFile := ""
	flag.StringVar(&outputFile, "output", "out.wav", "[[o]utput] file for the program")
	flag.StringVar(&outputFile, "o", "out.wav", "[[o]utput] file for the program")
	flag.Parse()
	nsamps := Duration * SampleRate
	angle := (math.Pi * 2) / float64(nsamps)
	egFile, err := os.Open("fixtures_32bit.wav")
	if err != nil {
		log.Fatalln(err)
	}
	egDecoder := wav.NewDecoder(egFile)
	egBuf, err := egDecoder.FullPCMBuffer()
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Open(outputFile)
	defer f.Close()
	if err == nil {
		log.Fatalln(outputFile + " already exists use something else")
	}
	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatalln(err)
	}
	outEncoder := wav.NewEncoder(
		outFile,
		SampleRate, // egBuf.Format.SampleRate,
		int(egDecoder.BitDepth),
		1, // egBuf.Format.NumChannels,
		int(egDecoder.WavAudioFormat),
	)

	outData := make([]int, nsamps)
	for i := 0; i < nsamps; i++ {
		sample := math.Sin(angle * Frequency * float64(i))
		outData[i] = int(sample * math.MaxInt32)

		if outData[i] < math.MinInt32 || outData[i] > math.MaxInt32 {
			fmt.Println(sample, outData[i])
		}
		// var buf [8]byte

		// i checked mine is LittleEndian cpu
		// might have to change to BigEndian for other CPU
		// binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
		// fmt.Println(int(sample * 10e10))
		// if i > 100 {
		// 	break
		// }
		// nBytes, err := outEncoder.Write(buf[:])
		// if err != nil {
		// 	log.Println("error while writing to the file")
		// 	log.Println(err)
		// } else {
		// 	// don't really know what this \r did
		// 	// but it seems like it is printing on the same line somehow
		// 	fmt.Printf("\rWrote: %d bytes to the file", nBytes)
		// }
	}

	outBuffer := audio.IntBuffer{
		Format:         egBuf.Format,
		Data:           outData,
		SourceBitDepth: egBuf.SourceBitDepth,
	}

	err = outEncoder.Write(&outBuffer)
	if err != nil {
		log.Fatalln(err)
	}

	err = outEncoder.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// reopen to confirm things worked well
	// out, err := os.Open("out.wav")
	// if err != nil {
	// 	panic(err)
	// }
	// d2 := wav.NewDecoder(out)
	// d2.ReadInfo()
	// fmt.Println("New file ->", d2)
	// out.Close()
}

// i think this is the sample rate
// or in simple terms how many points in a single sound wave
const nsamps = 50

func generate() {
	// don't know tau means
	// but what i have known so far is
	// in a simple sine wave you have
	// a Pi of samples the sound goes up and down /\
	// when you have 2Pi of samples of sound
	// it goes /\
	//           \/
	// completeing the sine wave
	tau := math.Pi * 2
	var angle float64 = tau / nsamps
	// In the continuous (real) world,
	// for generating a sine wave for a sound is
	// a(t) = A sin(2πft) = A sin(2πfn/sr) more info -> https://www.cs.nmsu.edu/~rth/cs/computermusic/Simple%20sound%20generation.html

	for i := 0; i < nsamps; i++ {
		samp := math.Sin(angle * float64(i))
		fmt.Printf("%.8f\n", samp)
	}
}
