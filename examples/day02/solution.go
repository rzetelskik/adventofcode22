package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Move int

const (
	Rock Move = iota + 1
	Paper
	Scissors
)

var stringToMove = map[string]Move{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

type Strategy int

const (
	Lose Strategy = iota + 1
	Draw
	Win
)

var stringToStrategy = map[string]Strategy{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

func getMoveScore(move Move) int {
	switch move {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}

	return 0
}

func getWinningCounterpart(move Move) Move {
	switch move {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	}

	return Move(0)
}

func getLosingCounterpart(move Move) Move {
	switch move {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	}

	return Move(0)
}

func getTurnScore(opponentMove, move Move) int {
	if opponentMove == move {
		return 3
	} else if getWinningCounterpart(move) == opponentMove {
		return 0
	} else {
		return 6
	}
}

func GetScore(opponentMove, move Move) int {
	return getMoveScore(move) + getTurnScore(opponentMove, move)
}

func GetMoveByStrategy(opponentMove Move, strategy Strategy) Move {
	switch strategy {
	case Lose:
		return getLosingCounterpart(opponentMove)
	case Draw:
		return opponentMove
	case Win:
		return getWinningCounterpart(opponentMove)
	}

	return Move(0)
}

func SolveFirstPart(in io.Reader, out io.Writer) (int, error) {
	total := 0

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		l := strings.Split(scanner.Text(), " ")
		opponentMove := stringToMove[l[0]]
		move := stringToMove[l[1]]
		total += GetScore(opponentMove, move)
	}

	return total, nil
}

func SolveSecondPart(in io.Reader, out io.Writer) (int, error) {
	total := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), " ")
		opponentMove := stringToMove[l[0]]
		strategy := stringToStrategy[l[1]]
		move := GetMoveByStrategy(opponentMove, strategy)
		total += GetScore(opponentMove, move)
	}

	return total, nil
}
