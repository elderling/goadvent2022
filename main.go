package main

import "fmt"

func main() {
	fmt.Println("Merry Christmas!")
	fmt.Printf("Day 01a Solution: %v\n", Day01aSolution(ReadElfCalories("my_day01a_data.txt")))
	fmt.Printf("Day 01b Solution: %v\n", Day01bSolution(ReadElfCalories("my_day01a_data.txt")))
	fmt.Printf("Day 02a Solution: %v\n", Day02aSolution("my_day02a_data.txt"))
	fmt.Printf("Day 02b Solution: %v\n", Day02bSolution("my_day02a_data.txt"))
	fmt.Printf("Day 03a Solution: %v\n", Day03aSolution("my_day03a_data.txt"))
	fmt.Printf("Day 03b Solution: %v\n", Day03bSolution("my_day03a_data.txt"))
}
