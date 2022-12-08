package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	start int
	end   int
}

func StringToSection(s string) Section {
	split := strings.Split(s, "-")

	start := split[0]
	end := split[1]

	section := Section{}

	section.start, _ = strconv.Atoi(start)
	section.end, _ = strconv.Atoi(end)

	return section
}

func (section1 Section) FullyContains(section2 Section) bool {
	if section1.start <= section2.start && section1.end >= section2.end {
		return true
	} else {
		return false
	}
}

func Day04aSolution(filename string) int {
	f, err := os.Open(filename)

	if err != nil {
		panic("Bad filename")
	}

	input := bufio.NewScanner(f)

	total := 0
	for input.Scan() {
		line := input.Text()

		sections := strings.Split(line, ",")

		section1 := StringToSection(sections[0])
		section2 := StringToSection(sections[1])

		if section1.FullyContains(section2) || section2.FullyContains(section1) {
			total++
		}
	}

	return total
}
