package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	Duration   = 2
	SampleRate = 44100
	Frequency  = 440
)

func main() {
	nsamps := Duration * SampleRate
	angle := (math.Pi * 2) / float64(nsamps)
	file, err := os.Create("out.bin")
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < nsamps; i++ {
		sample := math.Sin(angle * Frequency * float64(i))
		var buf [8]byte

		// i checked mine is LittleEndian cpu
		// might have to change to BigEndian for other CPU
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
		nBytes, err := file.Write(buf[:])
		if err != nil {
			log.Println("error while writing to the file")
			log.Println(err)
		} else {
			// don't really know what this \r did
			// but it seems like it is printing on the same line somehow
			fmt.Printf("\rWrote: %d bytes to the file", nBytes)
		}

	}
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
