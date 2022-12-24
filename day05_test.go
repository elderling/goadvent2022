package main

import (
	"bufio"
	"strings"
	"testing"
)

const TEST_DATA_DAY_05a = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `

func TestScanStartingStacks(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(TEST_DATA_DAY_05a))

	theSlice := ScanStartingStack(scanner)

	if theSlice[0] != " 1   2   3 " {
		t.Error("First Line busted")
	}

	if theSlice[3] != "    [D]    " {
		t.Error("Last Line busted")
	}
}

func TestInitStackBottoms(t *testing.T) {
	s := "[Z] [M] [P]"

	stacks := InitStackBottoms(s)

	if stacks[0][0] != "Z" {
		t.Error("nope 0")
	}

	if stacks[1][0] != "M" {
		t.Error("nope 1")
	}

	if stacks[2][0] != "P" {
		t.Error("nope 2")
	}
}

func TestInitTheRestOfTheStacks(t *testing.T) {
	s := "[Z] [M] [P]"
	stacks := InitStackBottoms(s)

	theRest := []string{
		"[N] [C]",
		"    [D]",
	}

	InitTheRestOfTheStacks(stacks, theRest)

	if stacks[0][0] != "Z" {
		t.Error("Corruption!")
	}

	if stacks[0][1] != "N" {
		t.Error("Corruption!")
	}

	if stacks[1][1] != "C" {
		t.Error("Corruption!")
	}
}
