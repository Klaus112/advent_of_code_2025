package parse

import (
	"fmt"
	"strconv"
	"strings"
)

//------------------------------------------------------------------

type IDPair struct {
	Start uint
	End   uint
}

func MustParseIDPair(s string) IDPair {
	parts := strings.Split(s, "-")

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Errorf("failed parsing start of string '%s': %w", s, err))
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Errorf("failed parsing end of string '%s': %w", s, err))
	}

	return IDPair{
		Start: uint(start),
		End:   uint(end),
	}
}
