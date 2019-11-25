package leetcode005

import "testing"

type Case struct {
	String string
	Expected string
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ String: "babad", Expected: "bab", },
		Case{ String: "cbbd", Expected: "bb", },
		Case{ String: "", Expected: "", },
		Case{ String: "a", Expected: "a", },
		Case{ String: "ac", Expected: "a", },
	}

	for _, e := range cases {
		actual := longestPalindrome(e.String)
		if actual != e.Expected {
			t.Errorf("got: %v, actual: %v", actual, e.Expected)
		}

	}
}
