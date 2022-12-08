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

func TestRucksackItemPriority(t *testing.T) {

	if RuckSackItemPriority(byte('a')) != 1 {
		t.Error("Priority of 'a' is Broken")
	}

	if RuckSackItemPriority(byte('A')) != 27 {
		t.Error("Priority of 'A' is Broken")
	}

	if RuckSackItemPriority(byte('Z')) != 27+25 {
		t.Error("Priority of 'Z' is Broken")
	}
}

func TestDay03aSolution(t *testing.T) {
	if Day03aSolution("test_data_day03a.txt") != 157 {
		t.Error("Day3aSolution() is broken")
	}
}

func TestRucksocksToBadgeGroups(t *testing.T) {
	sack_strings := ReadRucksackStrings("test_data_day03a.txt")
	sacks := StringsToRucksacks(sack_strings)

	groups := RucksacksToBadgeGroups(sacks)

	if string(groups[0].Rucksacks[0].CompartmentA) != "vJrwpWtwJgWr" {
		t.Error("Group 0, Rucksacks[0] is broken")
	}
}

func TestBadgeGroupIntersection(t *testing.T) {
	sack_strings := ReadRucksackStrings("test_data_day03a.txt")
	sacks := StringsToRucksacks(sack_strings)

	groups := RucksacksToBadgeGroups(sacks)

	intersection := BadgeGroupIntersection(&groups[0])

	if intersection != byte('r') {
		t.Error("Intersection is busted")
	}

}

func TestDay03bSolution(t *testing.T) {
	solution := Day03bSolution("test_data_day03a.txt")

	if solution != 70 {
		t.Error("Day 03b solution is busted")
	}
}
