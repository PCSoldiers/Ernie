// Package dart ...
package dart

import (
	"fmt"
	"sort"
	"strings"
)

const (
	// Bullseye ...
	Bullseye = 25
	// DoubleBullseye ...
	DoubleBullseye = 50
	// SingleCodePrefix constant
	SingleCodePrefix = "S"
	// DoubleCodePrefix constant
	DoubleCodePrefix = "D"
	// TripleCodePrefix constant
	TripleCodePrefix = "T"
	// SingleBullCodePrefix constant
	SingleBullCodePrefix = "SB"
	// DoubleBullCodePrefix constant
	DoubleBullCodePrefix = "DB"
)

// Score represents a darts Code and Value
type Score struct {
	code  string
	value int
}

// String adds ToString functionality to Score
func (s Score) String() string {
	return fmt.Sprintf("%3d = %s", s.value, s.code)
}

// GetScores ...
func GetScores() []Score {
	a := []Score{}
	a = append(a, getSingles()...)
	a = append(a, getDoubles()...)
	a = append(a, getTriples()...)
	a = append(a, getBulls()...)
	return a
}

// SortScores ...
func SortScores(s []Score) []Score {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].value < s[j].value
	})
	return s
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
func newScore(value int, code string) Score {
	var minScore = 1
	var maxScore = 60
	if value < minScore || value > maxScore {
		panic(fmt.Sprintf("Invalid value: %d. Valid range is %d to %d", value, minScore, maxScore))
	}
	return Score{code: code, value: value}
}
func getSingles() []Score {
	return getScoresUsing(makeSingle)
}
func getDoubles() []Score {
	return getScoresUsing(makeDouble)
}
func getTriples() []Score {
	return getScoresUsing(makeTriple)
}
func getBulls() []Score {
	return []Score{makeSingleBull(), makeDoubleBull()}
}
func getScoresUsing(fMakeScore func(int) Score) []Score {
	scores := make([]Score, 20)
	for i := range scores {
		value := i + 1
		scores[i] = fMakeScore(value)
	}
	return scores
}
func makeSingle(v int) Score {
	return newScore(toSingle(v), toCode(SingleCodePrefix, v))
}
func makeDouble(points int) Score {
	return newScore(toDouble(points), toCode(DoubleCodePrefix, points))
}
func makeTriple(v int) Score {
	return newScore(toTriple(v), toCode(TripleCodePrefix, v))
}
func makeSingleBull() Score {
	return newScore(Bullseye, toCode(SingleBullCodePrefix, Bullseye))
}
func makeDoubleBull() Score {
	return newScore(DoubleBullseye, toCode(DoubleBullCodePrefix, DoubleBullseye))
}
func toCode(p string, v int) string {
	if v == Bullseye || v == DoubleBullseye {
		return p	
	}
	return fmt.Sprintf("%s%d", p, v)
}
func toSingle(v int) int {
	return v
}
func toDouble(v int) int {
	return v * 2
}
func toTriple(v int) int {
	return v * 3
}
