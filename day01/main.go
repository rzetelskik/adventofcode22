package main

import (
	"bufio"
	"container/heap"
	_ "embed"
	"fmt"
	"golang.org/x/exp/constraints"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseInput(input string) ([][]int, error) {
	res := make([][]int, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))

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

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func solveFirstPart(input []int) int {
	res := 0

	for _, g := range input {
		res = max(res, g)
	}

	return res
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solveSecondPart(input []int) int {
	h := make(IntHeap, 0)

	for _, g := range input {
		heap.Push(&h, g)
		if h.Len() > 3 {
			heap.Pop(&h)
		}
	}

	res := 0
	for _, i := range h {
		res += i
	}

	return res
}

func main() {
	parsedInput, err := parseInput(input)
	if err != nil {
		fmt.Printf("can't parse input: %v\n", err)
		os.Exit(1)
	}
	summedInput := sumGroups(parsedInput)
	fmt.Println(solveFirstPart(summedInput))
	fmt.Println(solveSecondPart(summedInput))
}
