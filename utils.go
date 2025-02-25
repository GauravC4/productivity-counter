package main

import (
	"fmt"
	"strings"
	"time"
)

func beginOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func getProgressBar(hours float64, time string) string {
	maxValue := 12
	barLen := 4.0
	return fmt.Sprintf("[%s%s] %s / %dh\n", strings.Repeat("#", max(0, int(hours*barLen))), strings.Repeat(".", max(0, maxValue*int(barLen)-int(hours*barLen))), time, maxValue)
}
