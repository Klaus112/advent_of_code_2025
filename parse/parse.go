package parse

import "github.com/klaus112/advent_of_code_2025/files"

func InputLinebyLine[T any](path string, parseFn func(s string) T) []T {
	lines := files.ReadIntoSliceLineByLine(path)

	res := make([]T, 0, len(lines))
	for _, line := range lines {
		res = append(res, parseFn(line))
	}

	return res
}

func InputWithSeperator[T any](path string, seperator string, parseFn func(s string) T) []T {
	content := files.ReadWithSeperator(path, seperator)

	res := make([]T, 0, len(content))
	for _, line := range content {
		res = append(res, parseFn(line))
	}

	return res
}
