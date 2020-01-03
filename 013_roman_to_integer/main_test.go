package leetcode013

import "testing"

type Case struct {
	Input string
	Expected int
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Input: "III", Expected: 3, },
		Case{ Input: "IV", Expected: 4, },
		
	}

	for _, e := range cases {
		actual := romanToInt(e.Input)
		if actual != e.Expected {
			t.Errorf("Given : %v, got: %v, actual: %v", e.Input, actual, e.Expected)
		}

	}
}
