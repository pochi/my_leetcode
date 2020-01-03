package leetcode011

import "testing"

type Case struct {
	Input []int
	Expected int
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Input: []int{1,8,6,2,5,4,8,3,7}, Expected: 49, },
		Case{ Input: []int{1,1}, Expected: 1, },
		
	}

	for _, e := range cases {
		actual := maxArea(e.Input)
		if actual != e.Expected {
			t.Errorf("Given : %v, got: %v, actual: %v", e.Input, actual, e.Expected)
		}

	}
}
