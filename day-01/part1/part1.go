package part1

import (
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func Part1(words []string) int {
	num := 0

	for _, w := range words {
		for _, c := range w {
			if n, err := strconv.Atoi(string(c)); err == nil {
				num += 10 * n
				break
			}
		}

		for _, c := range reverse(w) {
			if n, err := strconv.Atoi(string(c)); err == nil {
				num += n
				break
			}
		}
	}

	return num
}
