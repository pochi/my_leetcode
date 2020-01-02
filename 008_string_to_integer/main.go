package leetcode008

import "math"
import "strconv"
import "regexp"
//import "strings"
//import "fmt"

func myAtoi(s string) int {
	max := int(math.Pow(2, 31))
	min := int(math.Pow(-2, 31))

	// space trim string
	// sts := strings.Replace(s, " ", "", -1)
	sts := regexp.MustCompile(`^\s+`).ReplaceAllString(s, "")
	n := string(regexp.MustCompile(`^[-+0-9][0-9]*`).Find([]byte(sts)))

	if n == "" || n == "+" || n == "-" {
		return 0
	}

	i, err := strconv.Atoi(n); if err != nil {
		switch err.(type) {
		case *strconv.NumError:
			if i > 0 {
				return max - 1
			}
			return min
		}
	}

	if i >= max {
		return max - 1
	}

	if i < min {
		return min
	}

	return i
}