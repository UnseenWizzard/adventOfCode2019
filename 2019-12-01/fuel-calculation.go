package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseInputFile(filename string) (moduleMasses []int) {
	file, err := os.Open(filename)
	checkForInputError(err)

	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')

		if err == io.EOF {
			return moduleMasses
		}
		checkForInputError(err)

		line = strings.TrimSuffix(line, "\n")

		var mass int
		mass, err = strconv.Atoi(line)
		checkForInputError(err)

		moduleMasses = append(moduleMasses, mass)
	}
}

func checkForInputError(err error) {
	if err != nil {
		log.Fatalf("Could not read input file!\n err: %s", err)
	}
}

func CalculateTotalFuelRequirement(moduleMasses []int) (total int) {
	total = 0
	for _, mass := range moduleMasses {
		total += CalcuateFuelRequirement(mass)
	}
	return total
}

func CalcuateFuelRequirement(mass int) int {
	return mass/3 - 2
}

// func main() {
// 	moduleMasses := ParseInputFile("inputdata.txt")
// 	total := calculateTotalFuelRequirement(moduleMasses)
// 	log.Printf("%d modules. Total Fuel Requirement: %d", len(moduleMasses), total)
// }