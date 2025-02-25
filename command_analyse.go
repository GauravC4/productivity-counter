package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func commandAnalyse(args []string) error {
	log.Printf("command analyse invloked with args %v\n", args)

	numberOfDays := 1
	verbose := false

	if 0 < len(args) {
		var err error
		numberOfDays, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid number of days")
		}
		if numberOfDays < 0 {
			numberOfDays *= -1
		}
	}

	if 1 < len(args) {
		if strings.ToLower(args[1]) == "v" {
			verbose = true
		}
	}

	bodToday := beginOfDay(time.Now())
	start := bodToday.Add(time.Hour * -24 * time.Duration(numberOfDays))
	end := time.Now()
	dailyTotals := make(map[time.Time]time.Duration)

	log.Printf("command analyse window start : %v, end : %v, verbose : %v\n", start, end, verbose)

	records := getRecordsFromWAL()

	if len(records) < 1 {
		fmt.Println("did not find any records in wal")
		return nil
	}

	i := len(records) - 1
	for i >= 0 {
		//log.Printf("compairing %v and %v, ans is : %v", start, records[i].startTime, start.Before(records[i].startTime))
		if start.After(records[i].startTime) {
			break
		}

		records[i].duration = records[i].endTime.Sub(records[i].startTime)

		// adding to daily total
		bod := beginOfDay(records[i].startTime)
		val, ok := dailyTotals[bod]
		if !ok {
			dailyTotals[bod] = time.Duration(0)
		}
		dailyTotals[bod] = val + records[i].duration

		i--
	}
	i++
	if i == len(records) {
		fmt.Println("did not find any records matching the given window")
		return nil
	}

	currDate := beginOfDay(records[i].startTime)
	for j := i; j < len(records); j++ {
		t := records[j].startTime
		et := records[j].duration

		// daily totals
		bod := beginOfDay(t)
		if !currDate.Equal(bod) {
			etTotal := dailyTotals[currDate]
			fmt.Printf("%v %s\n", currDate.Format("2006-01-02"), getProgressBar(etTotal.Hours(), etTotal.Truncate(time.Second).String()))
			currDate = bod
		}

		if verbose {
			fmt.Printf("%v %s", t.Format("03:04:05PM"), getProgressBar(et.Hours(), et.Truncate(time.Second).String()))
		}
	}

	// last days total not covered in loop
	etTotal := dailyTotals[currDate]
	fmt.Printf("%v %s\n", currDate.Format("2006-01-02"), getProgressBar(etTotal.Hours(), etTotal.Truncate(time.Second).String()))

	return nil
}
