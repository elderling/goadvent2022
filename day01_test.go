package main

import (
	"testing"
)

func TestReadElfCalories(t *testing.T) {
	elf_calories := ReadElfCalories("test_data_day01a.txt")

	if elf_calories[0][0] != 1000 {
		t.Error("didn't populate slice of slices correctly.")
	}
}

func TestDay01aSolution(t *testing.T) {
	test_calories := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}

	if Day01aSolution(test_calories) != 24000 {
		t.Error("didn't create slice of slices correctly")
	}
}
