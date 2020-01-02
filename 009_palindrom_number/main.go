package leetcode009

import "strconv"
//import "fmt"

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}

	if n < 10 {
		return true
	}

	s := strconv.Itoa(n)
	var (
		h int
		hf string
		hl string
	)	
	if len(s) % 2 == 0 {
		h = len(s) / 2
		hf, hl = s[0:h], s[h:]
	} else {
		h = (len(s) / 2) + 1
		hf, hl = s[0:h], s[h-1:]
	}

	hr := Reverse(hf)
	//fmt.Println(s, hf, hl, hr, hr == hl)
	
	if hr == hl {
		return true
	}	
	
	return false
}