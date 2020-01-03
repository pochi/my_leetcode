package leetcode012

import "strings"

func intToRoman(n int) string { 
	units := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	counts := make([]int, len(units))

	for i, u := range units {
		counts[i] = n / u
		n -= (n / u) * u
	}

	rs := ""
	for i, roman := range romans {
		rs += strings.Repeat(roman, counts[i])
	}

	return rs
}