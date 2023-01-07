package main

import (
	"errors"
	"fmt"
	"io"
	"math/bits"
	"strconv"
)

func findNUniqueEndPos(data []byte, n int) (int, error) {
	var bitset uint32
	for i := 0; i < n; i++ {
		bitset ^= 1 << uint32(data[i]-'a')
	}

	if bits.OnesCount32(bitset) == n {
		return n + 1, nil
	}

	for i := n; i < len(data); i++ {
		bitset ^= 1 << uint32(data[i]-'a')
		bitset ^= 1 << uint32(data[i-n]-'a')

		if bits.OnesCount32(bitset) == n {
			return i + 1, nil
		}
	}

	return 0, errors.New("can't find window of unique bytes")
}

func SolveFirstPart(in io.Reader, out io.Writer) error {
	data, err := io.ReadAll(in)
	if err != nil {
		return fmt.Errorf("can't read all from input: %w", err)
	}

	res, err := findNUniqueEndPos(data, 4)
	if err != nil {
		return err
	}

	_, err = out.Write([]byte(strconv.Itoa(res)))
	if err != nil {
		return fmt.Errorf("can't write to out: %w", err)
	}

	return nil
}

func SolveSecondPart(in io.Reader, out io.Writer) error {
	data, err := io.ReadAll(in)
	if err != nil {
		return fmt.Errorf("can't read all from input: %w", err)
	}

	res, err := findNUniqueEndPos(data, 14)
	if err != nil {
		return err
	}

	_, err = out.Write([]byte(strconv.Itoa(res)))
	if err != nil {
		return fmt.Errorf("can't write to out: %w", err)
	}

	return nil
}
