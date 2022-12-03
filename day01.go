package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func ReadElfCalories(filename string) [][]int {
	elf_calories := [][]int{}

	f, err := os.Open(filename)
	if err != nil {
		panic("couldn't open file")
	}

	input := bufio.NewScanner(f)

	single_elf := []int{}
	for input.Scan() {
		line := input.Text()

		if line == "" {
			elf_calories = append(elf_calories, single_elf)
			single_elf = []int{}
		} else {
			theInt, _ := strconv.Atoi(input.Text())
			single_elf = append(single_elf, theInt)
		}
	}

	return elf_calories
}

func Day01aSolution(calories [][]int) int {
	largest_total := 0

	for _, elf := range calories {
		subtotal := 0
		for _, cals := range elf {
			subtotal += cals
		}
		if subtotal > largest_total {
			largest_total = subtotal
		}
	}

	return largest_total
}

func CaloriesToTotalCalories(elf_calories [][]int) []int {
	totals := []int{}

	for _, cals := range elf_calories {
		total := 0
		for _, cal := range cals {
			total += cal
		}
		totals = append(totals, total)
	}

	return totals
}

func Day01bSolution(elf_calories [][]int) int {

	total_calories := CaloriesToTotalCalories(elf_calories)

	// we use greater than as the less than function to sort descending
	sort.Slice(total_calories, func(i, j int) bool { return total_calories[i] > total_calories[j] })

	return total_calories[0] + total_calories[1] + total_calories[2]
}
