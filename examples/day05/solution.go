package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/rzetelskik/adventofcode22/pkg/util"
	"io"
	"strconv"
	"strings"
)

const crateOffset = 4

type move struct {
	val  int
	from int
	to   int
}

func parseCrateStacks(in []string) []util.Stack[rune] {
	n := len(in[0])/crateOffset + 1

	stacks := make([]util.Stack[rune], n)
	for i := len(in) - 1; i >= 0; i-- {
		for si := range stacks {
			crate := rune(in[i][crateOffset*si+1])
			if crate != ' ' {
				stacks[si].Push(crate)
			}
		}
	}

	return stacks
}

func parseMoves(in []string) ([]move, error) {
	moves := make([]move, 0)

	for _, i := range in {
		ps := strings.Split(i, " ")

		val, err := strconv.Atoi(ps[1])
		if err != nil {
			return nil, err
		}

		from, err := strconv.Atoi(ps[3])
		if err != nil {
			return nil, err
		}

		to, err := strconv.Atoi(ps[5])
		if err != nil {
			return nil, err
		}

		moves = append(moves, move{val, from, to})
	}

	return moves, nil
}

func parseInput(in io.Reader) ([]util.Stack[rune], []move, error) {
	s := bufio.NewScanner(in)

	crateLines := make([]string, 0)
	for s.Scan() {
		if len(s.Text()) == 0 {
			break
		}
		crateLines = append(crateLines, s.Text())
	}

	moveLines := make([]string, 0)
	for s.Scan() {
		moveLines = append(moveLines, s.Text())
	}

	stacks := parseCrateStacks(crateLines)
	moves, err := parseMoves(moveLines)
	if err != nil {
		return nil, nil, fmt.Errorf("can't parse moves: %w", err)
	}

	return stacks, moves, nil
}

func SolveFirstPart(in io.Reader, out io.Writer) error {
	stacks, moves, err := parseInput(in)
	if err != nil {
		return fmt.Errorf("can't parse input: %w", err)
	}

	for _, m := range moves {
		for i := 0; i < m.val; i++ {
			stacks[m.to-1].Push(stacks[m.from-1].Top())
			stacks[m.from-1].Pop()
		}
	}

	var buf bytes.Buffer
	for _, s := range stacks {
		_, err := buf.WriteRune(s.Top())
		if err != nil {
			return fmt.Errorf("can't writeRune: %w", err)
		}
	}

	_, err = out.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("can't write to out: %w", err)
	}

	return nil
}

func SolveSecondPart(in io.Reader, out io.Writer) error {
	stacks, moves, err := parseInput(in)
	if err != nil {
		return fmt.Errorf("can't parse input: %w", err)
	}

	for _, m := range moves {
		var tmpStack util.Stack[rune]

		for i := 0; i < m.val; i++ {
			tmpStack.Push(stacks[m.from-1].Top())
			stacks[m.from-1].Pop()
		}
		for !tmpStack.IsEmpty() {
			stacks[m.to-1].Push(tmpStack.Top())
			tmpStack.Pop()
		}
	}

	var buf bytes.Buffer
	for _, s := range stacks {
		_, err := buf.WriteRune(s.Top())
		if err != nil {
			return fmt.Errorf("can't writeRune: %w", err)
		}
	}

	_, err = out.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("can't write to out: %w", err)
	}

	return nil
}
