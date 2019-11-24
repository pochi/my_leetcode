package leetcode003

import(
	"testing"
)

type Case struct {
	Input string
	Expected int
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Input: "abcabcbb", Expected: 3 },
		Case{ Input: "bbbbb", Expected: 1 },
		Case{ Input: "pwwkew", Expected: 3},
		Case{ Input: "", Expected: 0 },
		Case{ Input: "au", Expected: 2 },
		Case{ Input: "aab", Expected: 2 },
		Case{ Input: "bwf", Expected: 3 },		
	}

	for _, e := range cases {
		actual := lengthOfLongestSubstring(e.Input)
		if actual != e.Expected {
			t.Errorf("str: %v, got: %v, actual: %v", e.Input, actual, e.Expected)
		}

	}
}
