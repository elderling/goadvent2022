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
	OPPONENT = iota
	ME       = iota
	DRAW     = iota
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

func LookupOutcomeByLetter(letter string) int {
	if letter == "X" {
		return OPPONENT
	}
	if letter == "Y" {
		return DRAW
	}
	if letter == "Z" {
		return ME
	}

	panic("Invalid letter")
}

func RPSWinner(opponent, me int) int {
	lookup := [3][3]int{}
	lookup[ROCK][ROCK] = DRAW
	lookup[ROCK][PAPER] = ME
	lookup[ROCK][SCISORS] = OPPONENT

	lookup[PAPER][ROCK] = OPPONENT
	lookup[PAPER][PAPER] = DRAW
	lookup[PAPER][SCISORS] = ME

	lookup[SCISORS][ROCK] = ME
	lookup[SCISORS][PAPER] = OPPONENT
	lookup[SCISORS][SCISORS] = DRAW

	return lookup[opponent][me]
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
		opponent_shape := LookupRPSByLetter(letters[0])

		me_shape := LookupRPSByLetter(letters[1])

		total += RPSShapeScore(me_shape)

		winner := RPSWinner(opponent_shape, me_shape)

		if winner == ME {
			total += 6
		}
		if winner == OPPONENT {
			total += 0
		}
		if winner == DRAW {
			total += 3
		}
	}

	return total
}

func ScoreRPSDay02b(games []string) int {
	total := 0

	for _, game := range games {
		letters := strings.Split(game, " ")

		opponent_shape := LookupRPSByLetter(letters[0])

		me_shape := SelectRPSShapeForOutcome(opponent_shape, LookupOutcomeByLetter(letters[1]))

		winner := RPSWinner(opponent_shape, me_shape)

		total += RPSShapeScore(me_shape)

		if winner == ME {
			total += 6
		}
		if winner == OPPONENT {
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

func Day02bSolution(filename string) int {
	games := ReadRPSGames(filename)

	return ScoreRPSDay02b(games)
}
func SelectRPSShapeForOutcome(opponent_shape, outcome int) int {
	if opponent_shape == ROCK && outcome == DRAW {
		return ROCK
	}
	if opponent_shape == ROCK && outcome == ME {
		return PAPER
	}
	if opponent_shape == ROCK && outcome == OPPONENT {
		return SCISORS
	}
	if opponent_shape == PAPER && outcome == DRAW {
		return PAPER
	}
	if opponent_shape == PAPER && outcome == ME {
		return SCISORS
	}
	if opponent_shape == PAPER && outcome == OPPONENT {
		return ROCK
	}
	if opponent_shape == SCISORS && outcome == DRAW {
		return SCISORS
	}
	if opponent_shape == SCISORS && outcome == ME {
		return ROCK
	}
	if opponent_shape == SCISORS && outcome == OPPONENT {
		return PAPER
	}

	panic("Invalid shape or outcome")
}
