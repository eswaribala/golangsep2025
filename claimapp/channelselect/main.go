package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {

	// create channel
	bpChannel := make(chan int)
	ecgChannel := make(chan string)
	logChannel := make(chan string)
	alertChannel := make(chan string)
	quitChannel := make(chan bool)

	ecgMessage := []string{"Normal", "Abnormal", "Critical", "Warning", "Stable"}

	//bp producer
	go func() {
		//time delay
		t := time.NewTicker(800 * time.Millisecond)
		defer t.Stop()
		for {

			for range t.C {
				bpChannel <- gofakeit.IntRange(60, 230)
			}
		}
	}()
	//ecg producer
	go func() {
		//time delay
		t := time.NewTicker(800 * time.Millisecond)
		defer t.Stop()
		for {

			for range t.C {

				ecgChannel <- ecgMessage[gofakeit.IntRange(0, len(ecgMessage)-1)]
			}
		}
	}()

	//log producer
	go func() {
		//time delay
		t := time.NewTicker(2 * time.Second)
		defer t.Stop()
		for {

			for range t.C {
				logChannel <- gofakeit.HipsterSentence(10)
			}
		}
	}()

	//alert producer
	go func() {
		//time delay
		t := time.NewTicker(5 * time.Second)
		defer t.Stop()
		for {

			for range t.C {
				alertChannel <- "Alert: " + gofakeit.HipsterSentence(5)
			}
		}
	}()

	//quit producer
	go func() {
		fmt.Println("Press ENTER to quit...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		quitChannel <- true
	}()

	//receive channel
	for {
		select {
		case bp := <-bpChannel:
			fmt.Printf("Blood Pressure: %d mmHg\n", bp)
		case ecg := <-ecgChannel:
			fmt.Printf("ECG Status: %s\n", ecg)
		case log := <-logChannel:
			fmt.Printf("Log: %s\n", log)
		case alert := <-alertChannel:
			fmt.Printf("%s\n", alert)
		case <-quitChannel:
			fmt.Println("Quitting...")
			return
		}
	}

}
