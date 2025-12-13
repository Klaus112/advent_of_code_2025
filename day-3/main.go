package main

import (
	"fmt"
	"strconv"

	"github.com/klaus112/advent_of_code_2025/files"
)

func main() {
	batterieBanks := files.ReadIntoSliceLineByLine(files.DefaultFilePath)

	var sum uint
	for _, bank := range batterieBanks {
		// parse to ints
		var values []uint
		for _, r := range bank {
			val, _ := strconv.Atoi(string(r))
			values = append(values, uint(val))
		}

		// get firstSpot.
		highIdx, firstSpot := findHighest(values)

		var secondSpot uint
		if highIdx != len(values)-1 {
			_, secondSpot = findHighest(values[highIdx+1:])
		} else {
			// special case, the last spot has the highest value
			// so it automatically switches to the second spot.
			secondSpot = firstSpot
			// now we search for the first spot and exclude the last val.
			_, firstSpot = findHighest(values[:len(values)-1])
		}

		combinedSpots, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstSpot, secondSpot))
		sum += uint(combinedSpots)
	}

	fmt.Printf("Part1: Highest joltage %d\n", sum)
}

// findHighest returns the index and value of the highest uint.
func findHighest(values []uint) (int, uint) {
	// get highest.
	var (
		highest uint
		idx     int
	)
	for i, val := range values {
		if val > highest {
			idx = i
			highest = val
		}
	}

	return idx, highest
}
