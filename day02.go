package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	ROCK    = iota
	PAPER   = iota
	SCISORS = iota
)

const (
	PLAYER1 = iota
	PLAYER2 = iota
	DRAW    = iota
)

func LookupRPSByLetter(letter string) int {
	if letter == "A" || letter == "X" {
		return ROCK
	}
	if letter == "B" || letter == "Y" {
		return PAPER
	}
	if letter == "C" || letter == "Z" {
		return SCISORS
	}

	panic("Invalid letter")
}

func RPSWinner(player1, player2 int) int {
	lookup := [3][3]int{}
	lookup[ROCK][ROCK] = DRAW
	lookup[ROCK][PAPER] = PLAYER2
	lookup[ROCK][SCISORS] = PLAYER1

	lookup[PAPER][ROCK] = PLAYER1
	lookup[PAPER][PAPER] = DRAW
	lookup[PAPER][SCISORS] = PLAYER2

	lookup[SCISORS][ROCK] = PLAYER2
	lookup[SCISORS][PAPER] = PLAYER1
	lookup[SCISORS][SCISORS] = DRAW

	return lookup[player1][player2]
}

func ReadRPSGames(filename string) []string {
	f, err := os.Open(filename)

	if err != nil {
		panic("Couldn't open file.")
	}

	input := bufio.NewScanner(f)

	games := []string{}
	for input.Scan() {
		games = append(games, input.Text())
	}

	return games
}

func ScoreRPSGames(games []string) int {

	total := 0

	for _, game := range games {
		letters := strings.Split(game, " ")
		player1 := LookupRPSByLetter(letters[0])

		player2 := LookupRPSByLetter(letters[1])

		total += RPSShapeScore(player2)

		winner := RPSWinner(player1, player2)

		if winner == PLAYER2 {
			total += 6
		}
		if winner == PLAYER1 {
			total += 0
		}
		if winner == DRAW {
			total += 3
		}
	}

	return total
}

func RPSShapeScore(shape int) int {
	if shape == ROCK {
		return 1
	}
	if shape == PAPER {
		return 2
	}
	if shape == SCISORS {
		return 3
	}

	return -1
}

func Day02aSolution(filename string) int {
	games := ReadRPSGames(filename)

	return ScoreRPSGames(games)
}
