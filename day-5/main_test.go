package main

import (
	"testing"

	"github.com/klaus112/advent_of_code_2025/parse"
	"github.com/stretchr/testify/assert"
)

func Test_part2(t *testing.T) {
	tests := map[string]struct {
		// Named input parameters for target function.
		pairs []parse.IDPair
		want  uint
	}{
		"overlapping where end == start": {
			pairs: []parse.IDPair{
				{Start: 5, End: 7},
				{Start: 7, End: 8},
			},
			want: 4,
		},
		"none overlapping": {
			pairs: []parse.IDPair{
				{Start: 3, End: 5},
				{Start: 7, End: 8},
			},
			want: 5,
		},
		"multiple overlaps": {
			pairs: []parse.IDPair{
				{Start: 3, End: 8},
				{Start: 4, End: 7},
				{Start: 7, End: 10},
			},
			want: 8,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := part2(tt.pairs)
			assert.Equal(t, tt.want, got)
		})
	}
}
