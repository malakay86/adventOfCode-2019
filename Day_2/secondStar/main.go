package main

import (
	"fmt"
	"os"
)

func opCode_1(index int, input []int) {
	position1 := input[index+1]
	position2 := input[index+2]
	storePosition := input[index+3]

	input[storePosition] = input[position1] + input[position2]
}

func opCode_2(index int, input []int) {
	position1 := input[index+1]
	position2 := input[index+2]
	storePosition := input[index+3]

	input[storePosition] = input[position1] * input[position2]
}

func main() {
	input := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 19, 6, 23, 2, 23, 13, 27, 1, 27, 5, 31, 2, 31, 10, 35, 1, 9, 35, 39, 1, 39, 9, 43, 2, 9, 43, 47, 1, 5, 47, 51, 2, 13, 51, 55, 1, 55, 9, 59, 2, 6, 59, 63, 1, 63, 5, 67, 1, 10, 67, 71, 1, 71, 10, 75, 2, 75, 13, 79, 2, 79, 13, 83, 1, 5, 83, 87, 1, 87, 6, 91, 2, 91, 13, 95, 1, 5, 95, 99, 1, 99, 2, 103, 1, 103, 6, 0, 99, 2, 14, 0, 0}

	// 1202 program alarm previous modifications
	input[1] = 12
	input[2] = 2

	for index := 0; index < len(input); index += 4 {
		opCode := input[index]

		switch opCode {
		case 1:
			opCode_1(index, input)
		case 2:
			opCode_2(index, input)
		case 99:
			fmt.Println("opCode 99 found! Finishing program...")
			fmt.Printf("Result: %d\n", input[0])
			os.Exit(0)
		default:
			panic("Unexpected opCode")
		}
	}
}
