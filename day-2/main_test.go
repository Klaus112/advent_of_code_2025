package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countInvalidIDsInRange(t *testing.T) {
	tests := map[string]struct {
		in   idPair
		want uint
	}{
		"11-22": {
			in:   idPair{11, 22},
			want: 33,
		},
		"998-1012": {
			in:   idPair{998, 1012},
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
		in   idPair
		want uint
	}{
		"95-115": {
			in:   idPair{95, 115},
			want: 210,
		},
		"998-1012": {
			in:   idPair{1010, 1010},
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
