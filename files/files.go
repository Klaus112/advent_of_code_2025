package files

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

const (
	DefaultFilePath     = "input.txt"
	DefaultTestFilePath = "test-input.txt"
)

// ReadIntoSliceLineByLine reads the file at path into a slice of strings (line by line).
func ReadIntoSliceLineByLine(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to read file at %s: %s", path, err)
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

func ReadWithSeperator(path string, seperator string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to read file at %s: %s", path, err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("failed to read content from file: %s", err)
	}

	return strings.Split(string(content), seperator)
}
