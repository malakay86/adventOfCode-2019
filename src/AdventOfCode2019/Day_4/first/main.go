package main

import (
	"fmt"
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
		if previousDigit <= digit {
			previousDigit = digit
		} else {
			return false
		}
	}

	return true
}

func DigitsNeverDecrease(num int64) bool {
	numAsStr := strconv.FormatInt(num, 10)
	var previousDigit int32 = -1

	for _, digit := range numAsStr {
		if previousDigit <= digit {
			previousDigit = digit
		} else {
			return false
		}
	}
	return true
}

func main() {
	min := 387638
	max := 919123

	count := 0
	for i := min; i <= max; i++ {
		if SixDigitNumber(i) &&
			HasTwoRepeatedAdjacentDigits(int64(i)) &&
			DigitsNeverDecrease(int64(i)) {
			count++
		}
	}

	// Counter minus one is the valid result,
	// as the final i++ of the previous loop adds one more
	count--

	fmt.Println(count)
}
