package leetcode010

import "testing"

type Case struct {
	Input string
	Regexp string
	Expected bool
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Input: "aa", Regexp: "a", Expected: false },
		Case{ Input: "mississippi", Regexp: "mis*is*p*.", Expected: false },
		Case{ Input: "aab", Regexp: "c*a*b", Expected: true },
		Case{ Input: "ab", Regexp: ".*", Expected: true },
		Case{ Input: "aaa", Regexp: ".a", Expected: false },
		Case{ Input: "aaa", Regexp: "a.a", Expected: true },
		Case{ Input: "", Regexp: ".*", Expected: true },
		Case{ Input: "b", Regexp: "aaa.", Expected: false },

	}

	for _, e := range cases {
		actual := isMatch(e.Input, e.Regexp)
		if actual != e.Expected {
			t.Errorf("Given str: %v, regexp: %v, got: %v, actual: %v", e.Input, e.Regexp, actual, e.Expected)
		}

	}
}
