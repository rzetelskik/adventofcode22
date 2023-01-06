package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"github.com/rzetelskik/adventofcode22/pkg/util"
	"io"
	"strconv"
)

func parseInput(in io.Reader) ([][]int, error) {
	res := make([][]int, 0)

	scanner := bufio.NewScanner(in)

	cur := make([]int, 0)
	for scanner.Scan() {
		read := scanner.Text()

		if len(read) == 0 {
			res = append(res, cur)
			cur = make([]int, 0)
			continue
		}

		num, err := strconv.Atoi(read)
		if err != nil {
			return nil, err
		}

		cur = append(cur, num)
	}

	if len(cur) > 0 {
		res = append(res, cur)
	}

	return res, nil
}

func sumGroups(input [][]int) []int {
	res := make([]int, 0)

	for _, g := range input {
		total := 0
		for _, i := range g {
			total += i
		}
		res = append(res, total)
	}

	return res
}

func SolveFirstPart(in io.Reader, out io.Writer) error {
	parsedIn, err := parseInput(in)
	if err != nil {
		return fmt.Errorf("can't parse input: %w", err)
	}

	groups := sumGroups(parsedIn)

	res := 0
	for _, g := range groups {
		res = util.Max(res, g)
	}

	_, err = out.Write([]byte(strconv.Itoa(res)))
	if err != nil {
		return fmt.Errorf("can't write to out: %w", err)
	}

	return nil
}

func SolveSecondPart(in io.Reader, out io.Writer) error {
	parsedIn, err := parseInput(in)
	if err != nil {
		return fmt.Errorf("can't parse input: %w", err)
	}

	h := make(util.Heap[int], 0)

	groups := sumGroups(parsedIn)

	for _, g := range groups {
		heap.Push(&h, g)
		if h.Len() > 3 {
			heap.Pop(&h)
		}
	}

	res := 0
	for _, i := range h {
		res += i
	}

	_, err = out.Write([]byte(strconv.Itoa(res)))
	if err != nil {
		return fmt.Errorf("can't write to out: %w", err)
	}

	return nil
}
