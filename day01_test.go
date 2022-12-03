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
