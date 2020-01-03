package leetcode013

import "regexp"

func romanToInt(s string) int {
	units := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	total := 0

	for i, u := range units {
		r := "^(" + romans[i] + ")+"
		total += len(regexp.MustCompile(r).FindString(s)) / len(romans[i]) * u
		s = regexp.MustCompile(r).ReplaceAllString(s, "")
	}

	return total
}