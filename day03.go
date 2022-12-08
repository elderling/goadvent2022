package main

import (
	"bufio"
	"os"
)

type Rucksack struct {
	CompartmentA []byte
	CompartmentB []byte
}

type BadgeGroup struct {
	Rucksacks []Rucksack
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

func RucksacksToBadgeGroups(sacks []Rucksack) []BadgeGroup {
	groups := []BadgeGroup{}

	for i := 0; i < len(sacks); i = i + 3 {
		group := BadgeGroup{}
		group.Rucksacks = sacks[i : i+3]
		groups = append(groups, group)
	}

	return groups
}

func StringsToRucksacks(rs_strings []string) []Rucksack {
	sacks := []Rucksack{}

	for _, v := range rs_strings {
		sacks = append(sacks, StringToRucksack(v))
	}

	return sacks
}

func BadgeGroupIntersection(sack *BadgeGroup) byte {
	duplicate := map[byte]int{}

	for _, sack := range sack.Rucksacks {
		combined := append(sack.CompartmentA, sack.CompartmentB...)

		cmap := map[byte]bool{}

		for _, v := range combined {
			cmap[v] = true
		}

		for k := range cmap {
			duplicate[k]++
		}
	}

	var retval = byte(0)
	for k, v := range duplicate {
		if v == 3 {
			retval = k
			break
		}
	}

	return retval
}

func Day03bSolution(filename string) int {
	sack_strings := ReadRucksackStrings(filename)
	sacks := StringsToRucksacks(sack_strings)

	groups := RucksacksToBadgeGroups(sacks)

	total := 0

	for _, group := range groups {
		total += RuckSackItemPriority(BadgeGroupIntersection(&group))
	}

	return total
}
