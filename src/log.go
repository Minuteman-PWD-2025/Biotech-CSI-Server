package main

import (
	"fmt"
	"time"
)

func log(args ...string) {
	output := ""
	currentTime := time.Now()

	for _, str := range args {
		output += str
	}

	fmt.Println(currentTime.Format("2006-01-02 15:04:05") + " - " + output)
}
