package main

import (
	"bufio"
	"os"
)

type Rucksack struct {
	CompartmentA []byte
	CompartmentB []byte
}

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

func StringToRucksack(s string) Rucksack {
	rs := Rucksack{}

	half := len(s) / 2

	cA := s[0:half]
	cB := s[half:]

	rs.CompartmentA = []byte(cA)
	rs.CompartmentB = []byte(cB)

	return rs
}
