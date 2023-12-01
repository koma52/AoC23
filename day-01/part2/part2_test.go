package part2

import (
	"testing"
)

func TestPart2(t *testing.T) {
	words := []string{
		"two1nine",
		"eighttwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	result := Part2(words)

	if result != 281 {
		t.Fatalf("Part2 output is %d but got: %d", 281, result)
	}
}
