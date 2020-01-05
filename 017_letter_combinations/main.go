package leetcode017

func letterCombinations(digits string) []string {
	digitsDict := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	letters := make([]string, len([]rune(digits)))
	for i, d := range []rune(digits) {
		letters[i] = digitsDict[string(d)]
	}

	c := []string{}
	for _, l := range letters {
		c = product(c, l) 
	}

	return c
}

func product(c []string, l string) []string {
	r := []string{}

	if len(c) == 0 {
		for _, d := range []rune(l) {
			r = append(r, string(d))
		}
		return r
	}

	for _, x := range c {
		for _, y := range []rune(l) {
			r = append(r, x + string(y))
		}
	}

	return r
}