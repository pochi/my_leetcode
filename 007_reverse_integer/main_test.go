package leetcode007

import "testing"

type Case struct {
	Num int
	Expected int
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Num: 123, Expected: 321 },
		Case{ Num: -123, Expected: -321 },	
		Case{ Num: 1534236469, Expected: 0 },	
	}

	for _, e := range cases {
		actual := reverse(e.Num)
		if actual != e.Expected {
			t.Errorf("got integer: %v, actual: %v", actual, e.Expected)
		}

	}
}
