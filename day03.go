package main

import (
	"bufio"
	"os"
)

func ReadRucksackStrings(filename string) []string {
	rucksacks := []string{}

	f, err := os.Open(filename)

	if err != nil {
		panic("Unable to read file")
	}

	input := bufio.NewScanner(f)

	for input.Scan() {
		rucksacks = append(rucksacks, input.Text())
	}

	return rucksacks

}
