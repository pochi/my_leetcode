package leetcode016

import "testing"

type Case struct {
	Input []int
	Target int
	Expected int
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Input: []int{-1,2,1,-4}, Target: 1, Expected: 2, },
	}

	for _, e := range cases {
		actual := threeSumClosest(e.Input, e.Target)

		if actual != e.Expected {
			t.Errorf("Given : %v, target: %v, got: %v, actual: %v", e.Input, e.Target, actual, e.Expected)
		}

	}
}
