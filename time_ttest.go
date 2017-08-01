// +build test

package main

import (
	"fmt"
	"time"

	"github.com/mixer/clock"
)

func printTime() {
	date, _ := time.Parse(time.UnixDate, "Sat Mar  7 11:12:39 PST 2015")
	fmt.Println(clock.NewMockClock(date).Now())
}
