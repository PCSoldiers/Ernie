package main

import (
	"dart"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\n=== All DartScores (Values and Codes) ===")
	for _, s := range dart.GetScores() {
		line := fmt.Sprintf("Score{\"%s\", %d},", s.GetCode(), s.GetValue())
		fmt.Println(line)
	}
	// All Code for points
	fmt.Println("=== All Codes for single dart scores 1 to 60 ===")
	for i := 1; i <= 60; i++ {
		fmt.Printf("%d -> %s\n", i, strings.Join(dart.GetCodes(i), ", "))
	}
}
