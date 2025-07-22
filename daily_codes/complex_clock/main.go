package main

import (
	"fmt"
	"time"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"os"
)

func playSound() {
	f, err := os.Open("alarm.mp3")
	if err != nil {
		fmt.Println("Could not open the file", err)
		return
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("Could not decode the file", err)
		return
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func displayTime() {
	for {
		now := time.Now()
		fmt.Printf("\r%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())
		time.Sleep(1 * time.Second)
		if now.Second() == 0 {
			go playSound()
		}
	}
}

func main() {
	fmt.Println("Starting the complex clock...")
	displayTime()
}
