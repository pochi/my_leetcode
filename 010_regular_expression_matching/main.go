package leetcode010

import "regexp"
import "strings"

func isMatch(s string, p string) bool {
	if strings.Contains(p, ".") == false && strings.Contains(p, "*") == false {
		return s == p
	}

/* 	i := 0
	if s != "" {
		i = findIndex(p, string(s[0]))
	}

	if i == -1 {
		return false
	}

	p = p[i:]
 */
	if strings.Contains(p, "*") == false && len(s) != len(p) {
		return false
	}

	p = "^" + p + "$"
	

	//isMatchAstarisk(s, p)
	//r := strings.Replace(p, "*", "+", -1)
	return regexp.MustCompile(p).Match([]byte(s))
}

func findIndex(s string, c string) int {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == c || string(runes[i]) == "." {
			return i
		}
	}

	return -1
}

