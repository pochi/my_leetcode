package leetcode009

import "testing"

type Case struct {
	Int int
	Expected bool
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Int: 121, Expected: true },
		Case{ Int: -121, Expected: false },
		Case{ Int: 10, Expected: false },
		Case{ Int: 11, Expected: true },
		
	}

	for _, e := range cases {
		actual := isPalindrome(e.Int)
		if actual != e.Expected {
			t.Errorf("Given int: %v, got integer: %v, actual: %v", e.Int, actual, e.Expected)
		}

	}
}
