package leetcode012

import "testing"

type Case struct {
	Input int
	Expected string
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Input: 3, Expected: "III", },
		Case{ Input: 4, Expected: "IV", },		
		Case{ Input: 58, Expected: "LVIII", },		
	}

	for _, e := range cases {
		actual := intToRoman(e.Input)
		if actual != e.Expected {
			t.Errorf("Given : %v, got: %v, actual: %v", e.Input, actual, e.Expected)
		}

	}
}
