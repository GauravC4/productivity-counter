package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type record struct {
	startTime time.Time
	endTime   time.Time
}

const WAL_FILE = "wal.csv"

func writeSessionToWAL(startTime time.Time, endTime time.Time) {
	log.Printf("writeSessionToWAL involked with params startTime: %v, endTime: %v\n", startTime, endTime)

	file, err := os.OpenFile(WAL_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%v,%v\n", startTime.Local().String(), endTime.Local().String()))
	if err != nil {
		log.Fatalf("Failed to write to file: %v\n", err)
	}

	log.Println("writeSessionToWAL done")
}

func getRecordsFromWAL() []record {
	log.Println("getRecordsFromWAL involked ...")

	records := make([]record, 0)

	file, err := os.Open(WAL_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			lineArr := strings.Split(line, ",")
			if len(lineArr) < 2 {
				log.Printf("Invalid line in WAL : %v\n", line)
				continue
			}
			layout := "2006-01-02 15:04:05.999999 -0700 MST"
			startTime, err := time.Parse(layout, lineArr[0])
			if err != nil {
				log.Printf("Error parsing time string from wal : %v\n", lineArr[0])
				continue
			}
			endTime, err := time.Parse(layout, lineArr[1])
			if err != nil {
				log.Printf("Error parsing time string from wal : %v\n", lineArr[0])
				continue
			}

			records = append(records, record{startTime: startTime, endTime: endTime})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("getRecordsFromWAL done, found %d valid records\n", len(records))
	return records
}

func dumpToWAL(content string) {
	log.Printf("dumpToWAL involked for content len %d\n", len(content))

	file, err := os.OpenFile(WAL_FILE, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		log.Fatalf("Failed to write to file: %v\n", err)
	}

	log.Printf("dumpToWAL done")
}

func compressWAL() {
	log.Println("compressWAL involked ...")

	sessionsMap := make(map[time.Time]time.Time)

	records := getRecordsFromWAL()
	log.Printf("compressWAL found %d valid records\n", len(records))

	for _, record := range records {
		sessionsMap[record.startTime] = record.endTime
	}
	if len(records) == len(sessionsMap) {
		return
	}

	stringEntries := make([]string, len(sessionsMap))
	i := 0
	for k, v := range sessionsMap {
		stringEntries[i] = fmt.Sprintf("%v,%v", k, v)
		i += 1
	}

	dumpToWAL(strings.Join(stringEntries, "\n") + "\n")

	log.Printf("compressWAL done from %d to %d\n", len(records), len(sessionsMap))
}
