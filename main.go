package main

import "fmt"

func main() {
	fmt.Println("Merry Christmas!")
	fmt.Printf("Day 01a Solution: %v\n", Day01aSolution(ReadElfCalories("my_day01a_data.txt")))
	fmt.Printf("Day 01b Solution: %v\n", Day01bSolution(ReadElfCalories("my_day01a_data.txt")))
}
