package main

import (
	"testing"

	"github.com/klaus112/advent_of_code_2025/shared"
	"github.com/stretchr/testify/assert"
)

func Test_countInvalidIDsInRange(t *testing.T) {
	tests := map[string]struct {
		in   shared.IDPair
		want uint
	}{
		"11-22": {
			in:   shared.IDPair{Start: 11, End: 22},
			want: 33,
		},
		"998-1012": {
			in:   shared.IDPair{Start: 998, End: 1012},
			want: 1010,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := addUpInvalidIDsPart1(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_addUpInvalidIDsPart2(t *testing.T) {
	tests := map[string]struct {
		in   shared.IDPair
		want uint
	}{
		"95-115": {
			in:   shared.IDPair{Start: 95, End: 115},
			want: 210,
		},
		"998-1012": {
			in:   shared.IDPair{Start: 1010, End: 1010},
			want: 1010,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := addUpInvalidIDsPart2(tt.in)
			assert.Equal(t, tt.want, got, "Wanted %d; got %d", tt.want, got)
		})
	}
}
