package files

import (
	"bufio"
	"log"
	"os"
)

const (
	DefaultFilePath     = "input.txt"
	DefaultTestFilePath = "test-input.txt"
)

// ReadIntoSlice reads the file at path into a slice of strings (line by line).
func ReadIntoSlice(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to read file at %s: %w", path, err)
	}
	defer f.Close()

	res := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}
