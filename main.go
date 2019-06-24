package main

import (
	"fmt"
	"strconv"
	"time"
)

type HwAccepted struct {
	Id    int
	Grade int
	Date  time.Time
}

func (hwAccepted HwAccepted) String() string {
	return hwAccepted.Date.Format("02-02-2006") + " accepted " + strconv.Itoa(hwAccepted.Id) + " " + strconv.Itoa(hwAccepted.Grade)
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
	Date    time.Time
}

func (hwSubmitted HwSubmitted) String() string {
	return hwSubmitted.Date.Format("02-02-2006") + " submitted " + strconv.Itoa(hwSubmitted.Id) + " " + hwSubmitted.Comment
}

type OtusEvent interface {
	String() string
}

func LogOtusEvent(e OtusEvent) {
	fmt.Println(e)
}

func main() {
	fmt.Println("Hello, playground")
	LogOtusEvent(HwAccepted{12, 13, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)})
	LogOtusEvent(HwSubmitted{12, "code", "comment", time.Date(1992, time.November, 10, 23, 0, 0, 0, time.UTC)})
}
