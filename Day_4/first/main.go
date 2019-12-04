package main

import (
	"strconv"
)

func NumberOfDigits(num int) int {
	if num < 10 {
		return 1
	} else {
		return 1 + NumberOfDigits(num/10)
	}
}

func SixDigitNumber(num int) bool {
	return NumberOfDigits(num) == 6
}

func HasTwoRepeatedAdjacentDigits(num int64) bool {
	numAsStr := strconv.FormatInt(num, 10)
	var previousDigit rune
	for _, digit := range numAsStr {
		if previousDigit == digit {
			return true
		}
		previousDigit = digit
	}

	return false
}

func main() {

}
