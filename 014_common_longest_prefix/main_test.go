package leetcode014

import "testing"

type Case struct {
	Input []string
	Expected string
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Input: []string{"flower","flow","flight"}, Expected: "fl", },
		Case{ Input: []string{"dog","racecar","car"}, Expected: "", },		
		Case{ Input: []string{"aca","cba"}, Expected: "", },		

	}

	for _, e := range cases {
		actual := longestCommonPrefix(e.Input)
		if actual != e.Expected {
			t.Errorf("Given : %v, got: %v, actual: %v", e.Input, actual, e.Expected)
		}

	}
}
