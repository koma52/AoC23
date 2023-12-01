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
		s_num := []rune{'a', 'a'}
		g_index := 200000
		for i := 0; i < len(w); i++ {
			if p, err := strconv.Atoi(string(w[i])); err == nil {
				s_num[0] = rune(p)
				g_index = i
				break
			}
		}
		index := 200000000
		m := 0
		for i := 0; i < len(num_str); i++ {
			if strings.Index(w, string(num_str[i])) < index &&
				strings.Index(w, string(num_str[i])) != -1 {
				index = strings.Index(w, string(num_str[i]))
				m = i
			}
		}
		if index < g_index {
			s_num[0] = rune(m)
		}

		g_index = -1
		for i := len(w) - 1; i >= 0; i-- {
			if p, err := strconv.Atoi(string(w[i])); err == nil {
				s_num[1] = rune(p)
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
			s_num[1] = rune(m)
		}

		num += int(s_num[0]) * 10
		num += int(s_num[1])
	}

	return num
}
