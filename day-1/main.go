package main

import (
	"fmt"
	"strconv"

	"github.com/klaus112/advent_of_code_2025/files"
)

type direction struct {
	value string
}

func newDirection(value string) direction {
	return direction{value: value}
}

var (
	directionLeft  direction = newDirection("left")
	directionRight direction = newDirection("right")
)

type puzzleInput struct {
	direction direction
	rotation  uint
}

func (p puzzleInput) String() string {
	return fmt.Sprintf("direction:%s; rotation: %d", p.direction.value, p.rotation)
}

func rotatePart1(p puzzleInput, currentValue uint) uint {
	rest := p.rotation % 100

	switch p.direction {
	case directionLeft:
		if rest > currentValue {
			// would produce an underflow
			newVal := 100 + (int(currentValue) - int(rest))
			return uint(newVal)
		}

		return currentValue - rest
	case directionRight:
		newVal := rest + currentValue
		if newVal >= 100 {
			return newVal - 100
		}

		return newVal
	}

	panic(fmt.Errorf("should never reach"))
}

// rotatePart2 returns
//  1. the amount of times a rotation passed or hit 0.
//  2. the new value where the rotation landed at.
func rotatePart2(p puzzleInput, currentValue uint) (uint, uint) {
	rotationAmount := p.rotation / 100
	rest := p.rotation % 100

	switch p.direction {
	case directionLeft:
		if currentValue == 0 {
			return rotationAmount, 100 - rest
		}
		if rest > currentValue {
			// would produce an underflow
			newVal := 100 + (int(currentValue) - int(rest))
			return rotationAmount + 1, uint(newVal)
		}

		return rotationAmount, currentValue - rest
	case directionRight:
		newVal := rest + currentValue
		if newVal == 100 {
			return rotationAmount, newVal - 100
		}
		if newVal >= 100 {
			return rotationAmount + 1, newVal - 100
		}

		return rotationAmount, newVal
	}

	panic(fmt.Errorf("should never reach"))
}

func mustParsePuzzleInput(row string) puzzleInput {
	direction := directionLeft
	if row[0] == 'R' {
		direction = directionRight
	}

	rotation, err := strconv.ParseUint(row[1:], 10, 64)
	if err != nil {
		panic(fmt.Errorf("invalid value in rotation: %s", row[1:]))
	}

	return puzzleInput{
		direction: direction,
		rotation:  uint(rotation),
	}
}

//------------------------------------------------------------------

func main() {
	content := files.ReadIntoSliceLineByLine(files.DefaultFilePath)

	var (
		currentRotationValue uint = 50
		zeroHitsCount        uint
	)
	for _, row := range content {
		puzzleInput := mustParsePuzzleInput(row)
		currentRotationValue = rotatePart1(puzzleInput, currentRotationValue)
		if currentRotationValue == 0 {
			zeroHitsCount++
		}
	}

	currentRotationValue = 50
	var totalZeroHitPassCount uint

	for _, row := range content {
		puzzleInput := mustParsePuzzleInput(row)
		zeroHitPassCount, newVal := rotatePart2(puzzleInput, currentRotationValue)
		currentRotationValue = newVal
		totalZeroHitPassCount += zeroHitPassCount
		if currentRotationValue == 0 {
			totalZeroHitPassCount++
		}
	}

	fmt.Printf("Part1: Zero hits count: %d\n", zeroHitsCount)
	fmt.Printf("Part2: Zero hits + pass count: %d\n", totalZeroHitPassCount)
}
