package leetcode014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	lcp := strs[0]

	for i := 1; i < len(strs); i++ {
		lcp = commonPrefix(lcp, strs[i])
		if lcp == "" {
			return lcp
		}
	}

	/* for i := 1; i < len(strs); i++ {
		str = strs[i]
		if len(lcp) < len(str) {
			l = len(lcp)
		} else {
			l = len(str)
		}	

		for i:=0; i<l; i++ {
			if lcp[i] == str[i] {
				c += string(lcp[i])
			}
		}

	}	 */

	return lcp
}

func commonPrefix(lcp string, str string) string {
/* 	cp := ""
	tcp := ""

	for i := 0; i < len(lcp); i++ {
		for j := 0; j < len(str); j++ {
			if lcp[i] == str[j] {
				tcp = getCommonPrefix(lcp[i:], str[j:])
				if len(tcp) > len(cp) {
					cp = tcp
				}
			}
		}
	}

	return cp */
	return getCommonPrefix(lcp, str)
}

func getCommonPrefix(lcp string, str string) string {
	l := 0
	c := ""
	if len(lcp) < len(str) {
		l = len(lcp)
	} else {
		l = len(str)
	}

	for i:=0; i<l; i++ {
		if lcp[i] == str[i] {
			c += string(lcp[i])
		} else {
			return c
		}
	}

	return c
}