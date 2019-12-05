package main

import "testing"

func TestHasSixDigits(t *testing.T) {
	a := 123456
	if !SixDigitNumber(a) {
		t.Fail()
	}
}

func TestHasNotSixDigits(t *testing.T) {
	a := 12
	if SixDigitNumber(a) {
		t.Fail()
	}
}

func isInRange(min, max, number int) bool {
	if number >= min && number <= max {
		return true
	} else {
		return false
	}
}

func TestIsInRange(t *testing.T) {
	a := 34
	minRange := 33
	maxRange := 35

	if !isInRange(minRange, maxRange, a) {
		t.Fail()
	}
}

func TestIsNotInRange(t *testing.T) {
	a := 32
	minRange := 33
	maxRange := 35

	if isInRange(minRange, maxRange, a) {
		t.Fail()
	}
}

func TestIsInRange_checkLimits(t *testing.T) {
	a := 33
	minRange := 33
	maxRange := 35

	if !isInRange(minRange, maxRange, a) {
		t.Fail()
	}
}

func TestTwoAdjacentDigitsAreTheSame_OK(t *testing.T) {
	var a int64 = 3445

	if !HasTwoRepeatedAdjacentDigits(a) {
		t.Fail()
	}
}

func TestTwoAdjacentDigitsAreTheSame_NOT_OK(t *testing.T) {
	var a int64 = 3456

	if HasTwoRepeatedAdjacentDigits(a) {
		t.Fail()
	}
}

func TestTwoAdjacentDigitsAreTheSame_NOT_OK_2(t *testing.T) {
	var a int64 = 3444456

	if HasTwoRepeatedAdjacentDigits(a) {
		t.Fail()
	}
}

func TestTwoAdjacentDigitsAreTheSame_OK_at_beginning(t *testing.T) {
	var a int64 = 3356

	if !HasTwoRepeatedAdjacentDigits(a) {
		t.Fail()
	}
}

func TestTwoAdjacentDigitsAreTheSame_OK_at_end(t *testing.T) {
	var a int64 = 3456455

	if !HasTwoRepeatedAdjacentDigits(a) {
		t.Fail()
	}
}

func TestDigitsNeverDecrease(t *testing.T) {
	inputOutput := []struct {
		input int64
		output bool
	}{
		{23456, true},
		{234456, true},
		{23256, false},
		{34561, false},
	}

	for _, value := range inputOutput {
		if DigitsNeverDecrease(value.input) != value.output {
			t.Errorf("Fail TestDigitsNeverDecrease with numeric value %d\n", value.input)
		}
	}
}
