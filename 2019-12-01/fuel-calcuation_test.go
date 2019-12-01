package main

import (
	"testing"
	"fmt"
)

func assertFuelRequirement(expected int, acutal int, t *testing.T) {
	if acutal != expected {
		message := fmt.Sprintf("Fuel Req. should have been %d but was %d", expected, acutal)
		t.Fatal(message)
	}
}
	
func TestExpected1(t *testing.T) {
	assertFuelRequirement(2, CalcuateFuelRequirement(12), t)
}

func TestExpected2(t *testing.T) {
	assertFuelRequirement(2, CalcuateFuelRequirement(14), t)
}

func TestExpected3(t *testing.T) {
	assertFuelRequirement(654, CalcuateFuelRequirement(1969), t)
}

func TestExpected4(t *testing.T) {
	assertFuelRequirement(33583, CalcuateFuelRequirement(100756), t)
}

func TestExpectedTotal(t *testing.T) {
	masses := []int{12,14,1969,100756}
	assertFuelRequirement(34241, CalculateTotalFuelRequirement(masses), t)
}