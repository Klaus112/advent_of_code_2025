package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/klaus112/advent_of_code_2025/files"
)

func main() {
	content := files.ReadWithSeperator(files.DefaultFilePath, ",")

	// part 1
	var sum uint
	for _, s := range content {
		pair := mustParseInput(s)
		increment := addUpInvalidIDsPart1(pair)
		sum += increment
	}

	fmt.Printf("Part1: Invalid ID Count: %d\n", sum)

	// part 2
	sum = 0
	for _, s := range content {
		pair := mustParseInput(s)
		increment := addUpInvalidIDsPart2(pair)
		sum += increment
	}

	fmt.Printf("Part2: Invalid ID Count: %d\n", sum)
}

//------------------------------------------------------------------

type idPair struct {
	start uint
	end   uint
}

func mustParseInput(s string) idPair {
	parts := strings.Split(s, "-")

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Errorf("failed parsing start of string '%s': %w", s, err))
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Errorf("failed parsing end of string '%s': %w", s, err))
	}

	return idPair{
		start: uint(start),
		end:   uint(end),
	}
}

func addUpInvalidIDsPart1(pair idPair) uint {
	var res uint
	for i := pair.start; i <= pair.end; i++ {
		sNum := fmt.Sprint(i)
		sLen := len(sNum)
		if sLen%2 == 0 {
			if sNum[:sLen/2] == sNum[sLen/2:] {
				res += i
			}
		}
	}

	return res
}

func addUpInvalidIDsPart2(pair idPair) uint {
	var res uint
	for i := pair.start; i <= pair.end; i++ {
		sNum := fmt.Sprint(i)
		var searchStr = string(sNum[0])
		// we start with the first digit and alway add one more digit to the search string
		for j := 1; j < len(sNum); j++ {
			if isRepeating(sNum, searchStr) {
				fmt.Printf("%d is repeating\n", i)
				res += i
				break
			}
			// add next digit to search string
			add := sNum[j : j+1]
			searchStr = fmt.Sprintf("%s%s", searchStr, add)
		}

	}

	return res
}

func isRepeating(s, pattern string) bool {
	patternLen := len(pattern)
	otherLen := len(s)
	if otherLen%patternLen != 0 {
		// they are not divisable so they can't be repeating.
		return false
	}

	maxIterations := otherLen / patternLen
	for i := range maxIterations {
		startOffset := patternLen * i
		if pattern != s[startOffset:startOffset+patternLen] {
			return false
		}
	}

	return true
}
