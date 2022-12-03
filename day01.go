package main

import (
	"bufio"
	"fmt"
	"os"
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
			fmt.Println("Blank!")
			elf_calories = append(elf_calories, single_elf)
			single_elf = []int{}
		} else {
			fmt.Println(line)
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
