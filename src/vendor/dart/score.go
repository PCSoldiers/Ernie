// Package dart ...
package dart

import (
	"fmt"
	"strings"
)

const (
	// DoubleCodePrefix constant
	DoubleCodePrefix = "D"
	// TripleCodePrefix constant
	TripleCodePrefix = "T"
)

// Score represents a darts Code and Value
type Score struct {
	code  string
	value int
}

// GetCode ...
func (s Score) GetCode() string {
	return s.code
}

// GetValue ...
func (s Score) GetValue() int {
	return s.value
}

// String adds ToString functionality to Score
func (s Score) String() string {
	return fmt.Sprintf("%3d = %s", s.value, s.code)
}

// GetCodes returns all codes for the point value
func GetCodes(points int) []string {
	var codes []string
	for _, score := range GetScores() {
		if points == score.value {
			codes = append(codes, score.code)
		}
	}
	return codes
}

// IsLeagueOut returns true if the Score is a double, else false
func IsLeagueOut(s Score) bool {
	return (IsDoubleOut(s) || isTripleOut(s))
}

// IsDoubleOut returns true if the Score is a double, else false
func IsDoubleOut(s Score) bool {
	return strings.HasPrefix(s.code, DoubleCodePrefix)
}
func isTripleOut(s Score) bool {
	return strings.HasPrefix(s.code, TripleCodePrefix)
}

// GetScores ...
func GetScores() []Score {
	return []Score{
		Score{"S1", 1},
		Score{"S2", 2},
		Score{"D1", 2},
		Score{"S3", 3},
		Score{"T1", 3},
		Score{"S4", 4},
		Score{"D2", 4},
		Score{"S5", 5},
		Score{"S6", 6},
		Score{"D3", 6},
		Score{"T2", 6},
		Score{"S7", 7},
		Score{"S8", 8},
		Score{"D4", 8},
		Score{"S9", 9},
		Score{"T3", 9},
		Score{"S10", 10},
		Score{"D5", 10},
		Score{"S11", 11},
		Score{"S12", 12},
		Score{"D6", 12},
		Score{"T4", 12},
		Score{"S13", 13},
		Score{"S14", 14},
		Score{"D7", 14},
		Score{"S15", 15},
		Score{"T5", 15},
		Score{"S16", 16},
		Score{"D8", 16},
		Score{"S17", 17},
		Score{"S18", 18},
		Score{"D9", 18},
		Score{"T6", 18},
		Score{"S19", 19},
		Score{"S20", 20},
		Score{"D10", 20},
		Score{"T7", 21},
		Score{"D11", 22},
		Score{"D12", 24},
		Score{"T8", 24},
		Score{"SB", 25},
		Score{"D13", 26},
		Score{"T9", 27},
		Score{"D14", 28},
		Score{"D15", 30},
		Score{"T10", 30},
		Score{"D16", 32},
		Score{"T11", 33},
		Score{"D17", 34},
		Score{"D18", 36},
		Score{"T12", 36},
		Score{"D19", 38},
		Score{"T13", 39},
		Score{"D20", 40},
		Score{"T14", 42},
		Score{"T15", 45},
		Score{"T16", 48},
		Score{"DB", 50},
		Score{"T17", 51},
		Score{"T18", 54},
		Score{"T19", 57},
		Score{"T20", 60},
	}
}
