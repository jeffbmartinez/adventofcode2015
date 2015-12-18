// puzzle at http://adventofcode.com/day/10

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

var input = "1113122113"
var timesToRepeat = 50
var maxPrintDigits = 60

func main() {

	s := input

	fmt.Printf("%v:\t%v - %v\n", 0, len(s), s)

	for i := 0; i < timesToRepeat; i++ {
		s = NextOutput(s)

		digitsToShow := s
		elipsis := ""
		if len(s) > maxPrintDigits {
			digitsToShow = digitsToShow[:maxPrintDigits]
			elipsis = "..."
		}

		fmt.Printf("%v:\t%v - %v%v\n", i+1, len(s), digitsToShow, elipsis)
	}
}

func NextOutput(s string) string {
	if len(s) == 0 {
		fmt.Println("Warning: Empty input")
		return ""
	}

	var buffer bytes.Buffer

	consecutiveCount := 0
	previousChar := s[0]
	var ch byte

	for i := range s {
		ch = s[i]

		if ch == previousChar {
			consecutiveCount++
		} else {
			buffer.WriteString(strconv.Itoa(consecutiveCount))
			buffer.WriteByte(previousChar)

			consecutiveCount = 1
		}

		previousChar = ch
	}

	buffer.WriteString(strconv.Itoa(consecutiveCount))
	buffer.WriteByte(ch)

	return buffer.String()
}
