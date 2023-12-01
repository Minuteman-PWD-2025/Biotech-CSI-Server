package main

import (
	"fmt"
	"time"
)

// log is a function that prints log messages with timestamps.
// If isError is true the log message is printed in red color.
// If isError is false the log message is printed in white color.
func log(isError bool, args ...string) {
	output := ""
	currentTime := time.Now()

	// ansi escape codes
	const redColor = "\033[31m"
	const resetColor = "\033[0m"

	dateFormat := "2006-01-02 15:04:05"
	if isError {
		// if its an error print the date in red
		fmt.Printf("%s%s%s - ", redColor, currentTime.Format(dateFormat), resetColor)
	} else {
		// if its not an error print the date in white
		fmt.Print(currentTime.Format(dateFormat) + " - ")
	}

	for _, str := range args {
		output += str + " "
	}
	fmt.Println(output)
}
