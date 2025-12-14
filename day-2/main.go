package main

import (
	"fmt"

	"github.com/klaus112/advent_of_code_2025/files"
	"github.com/klaus112/advent_of_code_2025/parse"
)

func main() {
	pairs := parse.InputWithSeperator(files.DefaultFilePath, ",", parse.MustParseIDPair)

	// part 1
	var sum uint
	for _, pair := range pairs {
		increment := addUpInvalidIDsPart1(pair)
		sum += increment
	}

	fmt.Printf("Part1: Invalid ID Count: %d\n", sum)

	// part 2
	sum = 0
	for _, pair := range pairs {
		increment := addUpInvalidIDsPart2(pair)
		sum += increment
	}

	fmt.Printf("Part2: Invalid ID Count: %d\n", sum)
}

//------------------------------------------------------------------

func addUpInvalidIDsPart1(pair parse.IDPair) uint {
	var res uint
	for i := pair.Start; i <= pair.End; i++ {
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

func addUpInvalidIDsPart2(pair parse.IDPair) uint {
	var res uint
	for i := pair.Start; i <= pair.End; i++ {
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
