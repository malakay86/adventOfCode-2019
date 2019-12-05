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
	result := false
	var previousDigit rune
	counter := 1
	for _, digit := range numAsStr {
		if previousDigit == digit {
			if counter > 2 {
				return false
			}

			counter++
			result = true
		} else {
			counter = 1
		}

		previousDigit = digit
	}

	return result
}

func DigitsNeverDecrease (num int64) bool {
	numAsStr := strconv.FormatInt(num, 10)
	var previousDigit int32 = -1;

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
	min := 387638;
	max := 919123;

	count := 0
	for i := min; i <= max; i++ {
		if SixDigitNumber(i) &&
			HasTwoRepeatedAdjacentDigits(int64(i)) &&
			DigitsNeverDecrease(int64(i)) {
			count++
		}
	}

	fmt.Print(count)
}
