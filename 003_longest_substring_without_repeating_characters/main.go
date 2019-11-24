package leetcode003

import "strings"
const minLen = 1

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	
	maxLen := minLen
	noRepeatLen := 0
	
	for i := 0; i < len(s) - 1; i++ {
		unique := true			
		for j := i + 1; j < len(s); j++ {
			if strings.Contains(s[i:j], s[j:j+1]) {
				unique = false
				if maxLen < (j - i) {
					maxLen = j - i
				}
				break
			}
		}
		
		if unique && noRepeatLen == 0 {
			noRepeatLen = len(s) - i
		}
	}

	if noRepeatLen > maxLen {
		return noRepeatLen
	}
	
	return maxLen
}
