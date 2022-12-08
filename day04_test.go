package main

import (
	"testing"
)

func TestStringToSection(t *testing.T) {
	section := StringToSection("2-4")

	if section.start != 2 || section.end != 4 {
		t.Error("StringToSection busted")
	}
}

func TestFullyContains(t *testing.T) {
	section1 := Section{
		start: 2,
		end:   4,
	}

	section2 := Section{
		start: 2,
		end:   3,
	}

	if !section1.FullyContains(section2) {
		t.Error("Broken Fully Contains")
	}

	if section2.FullyContains(section1) {
		t.Error("Broken fully contains")
	}
}

func TestDay04aSolution(t *testing.T) {
	sol := Day04aSolution("test_data_day04a.txt")

	if sol != 2 {
		t.Error("Broken Day04aSolution")
	}
}
