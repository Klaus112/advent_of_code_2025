package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/klaus112/advent_of_code_2025/files"
	"github.com/klaus112/advent_of_code_2025/parse"
)

func main() {
	lines := files.ReadIntoSliceLineByLine(files.DefaultFilePath)

	var freshPairs []parse.IDPair
	for idx, line := range lines {
		if line == "" {
			// remove the pairs and empty line
			lines = lines[idx+1:]
			break
		}
		freshPairs = append(freshPairs, parse.MustParseIDPair(line))
	}

	var availableIngredients []uint
	for _, line := range lines {
		ingredient, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Errorf("failed to parse line %s: %w", line, err))
		}
		availableIngredients = append(availableIngredients, uint(ingredient))
	}

	fmt.Printf("Part 1: Fresh Ingredient Count: %d\n", part1(freshPairs, availableIngredients))
	fmt.Printf("Part 2: Fresh Ingredient Count: %d\n", part2(freshPairs))
}

func part1(freshPairs []parse.IDPair, availableIngredients []uint) uint {
	var count uint
	for _, ingredient := range availableIngredients {
		for _, freshPair := range freshPairs {
			if freshPair.Start <= ingredient && ingredient <= freshPair.End {
				count++
				break
			}
		}
	}

	return count
}

func part2(pairs []parse.IDPair) uint {
	// sort by start
	slices.SortFunc(pairs, func(a, b parse.IDPair) int {
		return int(a.Start) - int(b.Start)
	})

	// correct overlaps
	newPairs := make([]parse.IDPair, 0, len(pairs))
	for i := 0; i < len(pairs)-1; i++ {
		if pairs[i].End >= pairs[i+1].Start {
			// assign the current start to the next value and do not add this value to the new list.
			pairs[i+1].Start = pairs[i].Start
		} else {
			newPairs = append(newPairs, pairs[i])
		}

		// also check if the current values end is bigger than the last values end.
		if pairs[i].End > pairs[i+1].End {
			pairs[i+1].End = pairs[i].End
		}
	}

	// missing the last pair so add it here:
	lastPair := pairs[len(pairs)-1]
	secondLastPair := pairs[len(pairs)-2]
	if lastPair.End < secondLastPair.End {
		lastPair.End = secondLastPair.End
	}
	if secondLastPair.End >= lastPair.Start {
		lastPair.Start = secondLastPair.Start
	}

	newPairs = append(newPairs, pairs[len(pairs)-1])

	// now sum up the individual pairs
	var res uint
	for _, pair := range newPairs {
		res += pair.End - pair.Start + 1
	}

	return res // 558454643755254 = 0x1fbe95aaa08f6
}
