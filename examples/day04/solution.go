package main

import (
	"bufio"
	"github.com/rzetelskik/adventofcode22/pkg/util"
	"golang.org/x/exp/constraints"
	"io"
	"strconv"
	"strings"
)

func parseRanges(s string) ([][]int, error) {
	pairs := strings.SplitN(s, ",", 2)
	ranges, err := util.MapSliceWithError(pairs, func(p string) ([]int, error) {
		return util.MapSliceWithError(strings.Split(p, "-"), func(s string) (int, error) {
			return strconv.Atoi(s)
		})
	})
	if err != nil {
		return [][]int{}, err
	}

	return ranges, nil
}

func rangesOverlapFully[T constraints.Integer | constraints.Float](xs, ys []T) bool {
	if len(xs) == 0 || len(ys) == 0 {
		return true
	}

	min := util.Min(xs[0], ys[0])
	max := util.Max(xs[len(xs)-1], ys[len(ys)-1])

	return xs[0] == min && xs[len(xs)-1] == max || ys[0] == min && ys[len(ys)-1] == max
}

func rangesOverlap[T constraints.Integer | constraints.Float](xs, ys []T) bool {
	if len(xs) == 0 || len(ys) == 0 {
		return true
	}

	min := util.Min(xs[0], ys[0])
	max := util.Max(xs[len(xs)-1], ys[len(ys)-1])

	return (xs[len(xs)-1]-xs[0]+1)+(ys[len(ys)-1]-ys[0]+1) > max-min+1
}

func SolveFirstPart(in io.Reader, out io.Writer) (int, error) {
	scanner := bufio.NewScanner(in)

	res := 0
	for scanner.Scan() {
		ranges, err := parseRanges(scanner.Text())
		if err != nil {
			return 0, err
		}

		if rangesOverlapFully(ranges[0], ranges[1]) {
			res += 1
		}
	}
	return res, nil
}

func SolveSecondPart(in io.Reader, out io.Writer) (int, error) {
	scanner := bufio.NewScanner(in)

	res := 0
	for scanner.Scan() {
		ranges, err := parseRanges(scanner.Text())
		if err != nil {
			return 0, err
		}

		if rangesOverlap(ranges[0], ranges[1]) {
			res += 1
		}
	}
	return res, nil
}
