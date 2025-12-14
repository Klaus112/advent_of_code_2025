package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/klaus112/advent_of_code_2025/files"
	"github.com/klaus112/advent_of_code_2025/parse"
)

func main() {
	batterieBanks := parse.InputLinebyLine(files.DefaultFilePath, mustParseBatterieBank)

	fmt.Printf("Part1: Highest joltage %d\n", sumHighestJoltageInAllBanks(batterieBanks, 2))
	fmt.Printf("Part2: Highest joltage %d\n", sumHighestJoltageInAllBanks(batterieBanks, 12))
}

func sumHighestJoltageInAllBanks(batterieBanks []batterieBank, maxBatterieCount int) uint {
	var sum uint
	for _, bank := range batterieBanks {
		sum += findHighestJoltageForBank(bank, maxBatterieCount)
	}

	return sum
}

// findHighest returns the index and value of the highest uint.
func findHighest(values []uint, maxPossible uint) (int, uint) {
	// get highest.
	var (
		highest uint
		idx     int
	)
	for i, val := range values {
		if val > highest {
			idx = i
			highest = val
			if highest == maxPossible {
				break
			}
		}
	}

	return idx, highest
}

type batterieBank struct {
	values []uint
}

func mustParseBatterieBank(s string) batterieBank {
	// parse to ints
	var values []uint
	for _, r := range s {
		val, err := strconv.Atoi(string(r))
		if err != nil {
			panic(fmt.Errorf("could not convert %s to int: %w", string(r), err))
		}
		values = append(values, uint(val))
	}

	return batterieBank{
		values: values,
	}
}

func findHighestJoltageForBank(bb batterieBank, maxBatterieCount int) uint {
	// highestPossibleValue is the maximum joltage of a single batterie.
	const highestPossibleValue = 9

	highIdx := -1
	// we only search over the first 0:-batterieSearchCount(2 or 12)
	var highest uint = 0
	res := make([]uint, 0, maxBatterieCount)
	i := 1
	for {
		upperBound := len(bb.values) - maxBatterieCount + i

		// the start point of our next search is always the highIdx of the previous search + 1
		startIdx := highIdx + 1

		highIdx, highest = findHighest(bb.values[startIdx:upperBound], highestPossibleValue)
		// add start startIdx to highIdx
		highIdx += startIdx

		res = append(res, highest)
		highest = 0
		maxPossible := len(bb.values)
		if upperBound == maxPossible {
			break
		}
		i++
	}

	// Convert the individual values to a single joltage (concat with no seperator and parse to int again)
	var strBuilder strings.Builder
	for _, val := range res {
		fmt.Fprint(&strBuilder, val)
	}
	iterationJoltage, _ := strconv.Atoi(strBuilder.String())

	return uint(iterationJoltage)
}
