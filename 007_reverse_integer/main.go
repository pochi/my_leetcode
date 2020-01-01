package leetcode007

import "math"
import "strconv"

func reverse(n int) int {
	rs := ""
	m := false
	last := 0
	s := strconv.Itoa(n)
	max_s := int(math.Pow(2, 31))
	min_s := int(math.Pow(-2, 31))

	if n > max_s || n < min_s {
		return 0
	}

	if string(s[0]) == "-" {
		m = true
		last = 1
	}

	for l:=len(s); l > last; l-- {
		rs += string(s[l-1])
	}

	i, err := strconv.Atoi(rs)

	if m == true {
		i *= -1
	}

	if err != nil {
		return 0
	}

	if i > max_s || i < min_s {
		return 0
	}

	return i
}
