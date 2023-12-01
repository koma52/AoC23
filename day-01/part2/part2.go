package part2

import (
	"strconv"
	"strings"
)

func Part2(words []string) int {
	num := 0
	num_str := [10]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for _, w := range words {
		digits := [2]int{}
		g_index := len(w)
		for i := 0; i < len(w); i++ {
			if d, err := strconv.Atoi(string(w[i])); err == nil {
				digits[0] = d
				g_index = i
				break
			}
		}
		index := len(w)
		m := 0
		for i := 0; i < len(num_str); i++ {
			if strings.Index(w, string(num_str[i])) < index &&
				strings.Index(w, string(num_str[i])) != -1 {
				index = strings.Index(w, string(num_str[i]))
				m = i
			}
		}
		if index < g_index {
			digits[0] = m
		}

		g_index = -1
		for i := len(w) - 1; i >= 0; i-- {
			if d, err := strconv.Atoi(string(w[i])); err == nil {
				digits[1] = d
				g_index = i
				break
			}
		}
		index = -1
		m = 0
		for i := 0; i < len(num_str); i++ {
			if strings.LastIndex(w, string(num_str[i])) > index &&
				strings.LastIndex(w, string(num_str[i])) != -1 {
				index = strings.LastIndex(w, string(num_str[i]))
				m = i
			}
		}
		if index > g_index {
			digits[1] = m
		}

		num += digits[0] * 10
		num += digits[1]
	}

	return num
}
