package main

import (
	"fmt"

	"github.com/klaus112/advent_of_code_2025/files"
	"github.com/klaus112/advent_of_code_2025/parse"
)

const (
	searchRune = "@"
	markedRune = "x"
)

func main() {
	board := parse.InputLinebyLine(files.DefaultFilePath, mustParseRow)

	for i := range board {
		for j := range board[i] {
			if board[i][j] == searchRune {
				// paper roll check the neighbours
				if hasFewerThanFourNeighboursFilled(board, i, j) {
					board[i][j] = markedRune
				}

			}
		}
	}

	printBoard(board)

	var markSum uint
	for i := range board {
		for j := range board[i] {
			if board[i][j] == markedRune {
				markSum++
			}
		}
	}
	fmt.Printf("Accessible paper count: %d\n", markSum)
}

func printBoard(board [][]string) {
	for i := range board {
		for j := range board[i] {
			fmt.Print(board[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("")
}

// hasFewerThanFourNeighboursFilled returns true if more than 3 neighbours have paper rolls.
func hasFewerThanFourNeighboursFilled(board [][]string, i, j int) bool {
	if i == 0 && j == 7 {
		fmt.Println("found it")
	}
	var counter uint

	// check below
	{
		newI := i + 1
		for _, newJ := range []int{j - 1, j, j + 1} {
			if isOutOfBounds(board, newI, newJ) {
				// skip impossible fields
				continue
			}

			if board[newI][newJ] == searchRune {
				counter++
			}
		}
	}

	// check above
	{
		newI := i - 1
		for _, newJ := range []int{j - 1, j, j + 1} {
			if isOutOfBounds(board, newI, newJ) {
				// skip impossible fields
				continue
			}

			if board[newI][newJ] == searchRune || board[newI][newJ] == markedRune {
				counter++
			}
		}
	}

	// check right
	{
		if !isOutOfBounds(board, i, j+1) && board[i][j+1] == searchRune {
			counter++
		}
	}
	// check left
	{
		if !isOutOfBounds(board, i, j-1) && (board[i][j-1] == searchRune || board[i][j-1] == markedRune) {
			counter++
		}
	}

	return counter < 4
}

func isOutOfBounds(board [][]string, i, j int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
		return true
	}

	return false
}

func mustParseRow(s string) []string {
	row := make([]string, 0, len(s))
	for _, r := range s {
		row = append(row, string(r))
	}

	return row
}
