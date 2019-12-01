package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateFuelForMass(mass int) (fuel int) {
	fuel = (mass / 3) - 2
	return
}

func calculateFuelForFuel(fuel int) int {
	var totalFuel, currentFuel = 0, fuel

	for currentFuel > 0 {
		fuelForFuel := calculateFuelForMass(currentFuel)
		if fuelForFuel > 0 {
			totalFuel += fuelForFuel
		}
		currentFuel = fuelForFuel
	}

	return totalFuel
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file")
	}
	defer file.Close()

	// Read line by line to perform calculation
	scanner := bufio.NewScanner(file)
	accumulated := 0
	for scanner.Scan() {
		if massAsInt, err := strconv.Atoi(scanner.Text()); err == nil {
			fuelNeededForMass := calculateFuelForMass(massAsInt)

			// Accumulate the fuel for the mass
			accumulated += fuelNeededForMass

			// Calculate fuel needed for the fuel itself
			if fuelNeededForMass > 0 {
				accumulated += calculateFuelForFuel(fuelNeededForMass)
			}
		}
	}

	fmt.Println("Total fuel amount needed:", accumulated)
}
