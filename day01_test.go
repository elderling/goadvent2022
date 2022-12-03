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

func TestCaloriesToTotalCalories(t *testing.T) {
	test_calories := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}

	total_cals := CaloriesToTotalCalories(test_calories)

	if total_cals[0] != 6000 {
		t.Error("didn't create total slice correctly")
	}

}

func TestDay01bSolution(t *testing.T) {
	test_calories := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}

	solution := Day01bSolution(test_calories)

	if solution != 45000 {
		t.Error("wrong Day 01b solution")
	}
}
