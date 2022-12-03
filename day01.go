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
