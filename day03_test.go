package main

import (
	"testing"
)

func TestReadRucksackStrings(t *testing.T) {
	rucksacks := ReadRucksackStrings("test_data_day03a.txt")

	if rucksacks[0] != "vJrwpWtwJgWrhcsFMMfFFhFp" {
		t.Error("First line not read properly")
	}

	if rucksacks[5] != "CrZsJsPPZsGzwwsLwLmpwMDw" {
		t.Error("Last line not read properly")
	}
}

func TestStringToRuckSack(t *testing.T) {
	s := "vJrwpWtwJgWrhcsFMMfFFhFp"

	rs := StringToRucksack(s)

	if string(rs.CompartmentA) != "vJrwpWtwJgWr" {
		t.Error("Compartment A is busted")
	}

	if string(rs.CompartmentB) != "hcsFMMfFFhFp" {
		t.Error("Compartment B is busted")
	}
}

func TestFindRucksackDuplicate(t *testing.T) {
	s := "vJrwpWtwJgWrhcsFMMfFFhFp"

	rs := StringToRucksack(s)

	duplicate := FindRucksackDuplicate(&rs)

	if string(duplicate) != "p" {
		t.Error("Duplicate broken")
	}
}
