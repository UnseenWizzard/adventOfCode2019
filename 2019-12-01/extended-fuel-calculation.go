package main

import (
	"log"
)

func ExtendedCalculateTotalFuelRequirement(moduleMasses []int) (total int) {
	total = 0
	for _, mass := range moduleMasses {
		currentMass := mass
		for {
			requiredFuel := CalcuateFuelRequirement(currentMass)
			if (requiredFuel < 0) {
				break
			}
			total += requiredFuel
			currentMass = requiredFuel
		}
	}
	return total
}

// func main() {
// 	moduleMasses := ParseInputFile("inputdata.txt")
// 	total := ExtendedCalculateTotalFuelRequirement(moduleMasses) 
// 	log.Printf("%d modules. Total Fuel Requirement: %d", len(moduleMasses), total)
// }