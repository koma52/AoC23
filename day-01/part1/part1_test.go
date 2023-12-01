package part1

import "testing"

func TestPart1(t *testing.T) {
	var words = []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}

	result := Part1(words)

	if result != 142 {
		t.Fatalf("Part1 output is %d but got %d", 142, result)
	}
}
