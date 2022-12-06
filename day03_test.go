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
