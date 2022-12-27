package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"log"
	"os"
	"time"
)
import "github.com/faiface/beep/speaker"

func main() {
	// open file
	f, err := os.Open("bg.mp3")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("open file success")

	// decoded
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	log.Println("decoded file success")

	// init speaker
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("speaker init success")

	// start play
	log.Println("start playing")
	loop := beep.Loop(-1, streamer)
	done := make(chan bool)
	speaker.Play(beep.Seq(loop, beep.Callback(func() {
		done <- true
	})))

	// wait finished
	for {
		select {
		case <-done:
			log.Println("play end")
			return
			//case <-time.After(time.Second):
			//	speaker.Lock()
			//	fmt.Println(format.SampleRate.D(streamer.Position()).Round(time.Second))
			//	speaker.Unlock()
		}
	}
}
