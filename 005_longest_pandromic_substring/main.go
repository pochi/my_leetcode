package leetcode005

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	longestString := ""
	
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			if len(longestString) >= len(s[i:j]) {
				break
			}

			// Speed optimization 
			if s[i:i+1] != s[j-1:j] {
				continue
			}
			
			if s[i:j] == reverse(s[i:j]) {
				if len(longestString) < len(s[i:j]) {
					longestString = s[i:j]
				}
			}
		}
	}

	return longestString
}
