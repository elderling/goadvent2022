package main

import (
	"bufio"
	"regexp"
)

func ScanStartingStack(scanner *bufio.Scanner) []string {
	s := []string{}

	var matchNumeric = regexp.MustCompile(` \d `)

	for scanner.Scan() {
		t := scanner.Text()

		s = append([]string{t}, s...)

		if matchNumeric.MatchString(t) {
			break
		}
	}

	return s
}

func InitStackBottoms(s string) [][]string {
	stacks := [][]string{}

	for _, v := range StackStringToSlice(s) {
		stack := []string{v}
		stacks = append(stacks, stack)
	}

	return stacks
}

func InitTheRestOfTheStacks(stacks [][]string, theRest []string) {
	for _, val := range theRest {
		spec := StackStringToSlice(val)

		for i, theVal := range spec {
			if theVal != " " {
				stacks[i] = append(stacks[i], theVal)
			}
		}
	}
}

func StackStringToSlice(s string) []string {
	stackString := []string{}

	for i := 1; i < len(s); i = i + 4 {
		stackString = append(stackString, string(s[i]))
	}

	return stackString
}

func MoveSingleCrate(stacks [][]string, from int, to int) {
	stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
	stacks[from] = stacks[from][:len(stacks[from])-1]
}

/*
func GetStack
func PushStack
func PopStack

*/
