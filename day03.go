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

func FindRucksackDuplicate(rucksack *Rucksack) byte {
	dupfinder := map[byte]int{}

	for _, v := range rucksack.CompartmentA {
		dupfinder[v]++
	}

	var return_val = byte(0)

	for _, v := range rucksack.CompartmentB {
		_, exists := dupfinder[v]

		if exists {
			return_val = v
			break
		}
	}

	return return_val
}

func RuckSackItemPriority(item byte) int {

	if item > 96 {
		return int(item - 96)
	} else {
		return int(item - 65 + 27)
	}
}

func Day03aSolution(filename string) int {
	rs_strings := ReadRucksackStrings(filename)

	total := 0

	for _, v := range rs_strings {
		rs := StringToRucksack(v)

		dup := FindRucksackDuplicate(&rs)

		priority := RuckSackItemPriority(dup)

		total += priority
	}

	return total
}
