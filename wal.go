package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func writeSessionToWAL(startTime time.Time, endTime time.Time) {
	log.Printf("writeSessionToWAL involked with params startTime: %v, endTime: %v", startTime, endTime)

	file, err := os.OpenFile("wal.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%v,%v\n", startTime.Local().String(), endTime.Local().String()))
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	log.Printf("writeSessionToWAL done")
}
