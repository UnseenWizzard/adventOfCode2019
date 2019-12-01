package main

import "testing"
	
func TestExtendedExpected1(t *testing.T) {
	masses := []int{12}
	assertFuelRequirement(2, ExtendedCalculateTotalFuelRequirement(masses), t)
}

func TestExtendedExpected2(t *testing.T) {
	masses := []int{14}
	assertFuelRequirement(2, ExtendedCalculateTotalFuelRequirement(masses), t)
}

func TestExtendedExpected3(t *testing.T) {
	masses := []int{1969}
	assertFuelRequirement(966, ExtendedCalculateTotalFuelRequirement(masses), t)
}

func TestExtendedExpected4(t *testing.T) {
	masses := []int{100756}
	assertFuelRequirement(50346, ExtendedCalculateTotalFuelRequirement(masses), t)
}

func TestExtendedExpectedTotal(t *testing.T) {
	masses := []int{12,14,1969,100756}
	assertFuelRequirement(51316, ExtendedCalculateTotalFuelRequirement(masses), t)
}