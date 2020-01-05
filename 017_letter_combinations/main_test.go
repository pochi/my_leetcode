package leetcode017

import "testing"
import "reflect"

type Case struct {
	Input string
	Expected []string
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Input: "23", Expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, },
	}

	for _, e := range cases {
		actual := letterCombinations(e.Input)

		if reflect.DeepEqual(actual, e.Expected) == false {
			t.Errorf("Given : %v, got: %v, actual: %v", e.Input, actual, e.Expected)
		}

	}
}
