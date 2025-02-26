package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const CHECKPOINT_INTERVAL = 60 * 5 // 5 mins

func commandStart(args []string) error {
	log.Printf("commandStart invoked ...")

	stop := make(chan bool) // Channel to signal stopping timer
	done := make(chan bool) // Channel to signal when everything is done
	go timer(stop, done)
	go listenToUserStop(stop)
	<-done
	log.Printf("commandStart done")
	return nil
}

func timer(stop <-chan bool, done chan<- bool) {
	log.Println("Timer function invoked ... ")
	startTime := time.Now()
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	checkpointTime := 0

	log.Printf("Starting session at : %v", startTime)
	for {
		select {
		case t := <-ticker.C:
			elapsedTime := t.Sub(startTime)
			fmt.Printf("\rElapsed time : %v", elapsedTime.Truncate(time.Second).String())

			// save to wal every checkpoint
			checkpointTime += 1
			if checkpointTime >= CHECKPOINT_INTERVAL {
				writeSessionToWAL(startTime, time.Now())
				checkpointTime = 0
			}
		case <-stop:
			done <- true

			// cleanup
			writeSessionToWAL(startTime, time.Now())
			compressWAL()
			log.Println("Timer function stopped")
			return
		}

	}

}

func listenToUserStop(stop chan<- bool) {
	log.Println("Listening for user input to stop timer ... ")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Press 'z' and Enter to stop")
		input, _ := reader.ReadString('\n')
		if strings.ToLower(input) == "z\n" {
			log.Println("User input received, stopping timer")
			stop <- true
			return
		}
	}

}
