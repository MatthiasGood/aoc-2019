package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	sum, err := sumOfFuelRequirements(reader); if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("the sum of the fuel requirements for all of the modules on the spacecraft is: %d\n", sum)
}

func sumOfFuelRequirements(reader *bufio.Reader) (int, error) {
	var sum int
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		mass, err := strconv.Atoi(string(line)); if err != nil {
			return sum, fmt.Errorf("line expected to be a number, but got: %s", line)
		}

		// Part 1
		// sum += findFuelForModule(mass)

		// Part 2
		sum += findFuelForModuleWithFuelMass(mass)
	}
	return sum, nil
}

func findFuelForModuleWithFuelMass(mass int) int {
	fuel := findFuelForModule(mass)
	if fuel <= 0 {
		return 0
	}
	return fuel + findFuelForModuleWithFuelMass(fuel)
}

/*
	To find the fuel required for a module, take its mass,
	divide by three, round down, and subtract 2.
 */
func findFuelForModule(mass int) int {
	return (mass / 3) - 2
}