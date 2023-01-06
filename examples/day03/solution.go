package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

var priorities = map[rune]int{}

func init() {
	for i := 0; i < 26; i++ {
		priorities[rune('a'+i)] = i + 1
		priorities[rune('A'+i)] = i + 27
	}
}

func stringToSet(s string) map[int]struct{} {
	set := map[int]struct{}{}

	for _, c := range s {
		set[priorities[rune(c)]] = struct{}{}
	}

	return set
}

func intersectSets[T comparable](xs, ys map[T]struct{}) map[T]struct{} {
	intersection := map[T]struct{}{}

	for x, _ := range xs {
		if v, ok := ys[x]; ok {
			intersection[x] = v
		}
	}

	return intersection
}

func SolveFirstPart(in io.Reader, out io.Writer) error {
	total := 0

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		firstHalf := stringToSet(line[0 : len(line)/2])
		secondHalf := stringToSet(line[len(line)/2 : len(line)])
		intersection := intersectSets(firstHalf, secondHalf)

		for i, _ := range intersection {
			total += i
		}
	}

	_, err := out.Write([]byte(strconv.Itoa(total)))
	if err != nil {
		return fmt.Errorf("can't write to out: %w", err)
	}

	return nil
}

func SolveSecondPart(in io.Reader, out io.Writer) error {
	total := 0

	lines := make([]string, 0)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i+3 <= len(lines); i += 3 {
		sets := make([]map[int]struct{}, 3)
		for j := 0; j < 3; j++ {
			sets[j] = stringToSet(lines[i+j])
		}
		intersection := intersectSets(intersectSets(sets[0], sets[1]), sets[2])
		for i, _ := range intersection {
			total += i
		}
	}

	_, err := out.Write([]byte(strconv.Itoa(total)))
	if err != nil {
		return fmt.Errorf("can't write to out: %w", err)
	}

	return nil
}
