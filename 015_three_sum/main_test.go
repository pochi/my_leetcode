package leetcode015

import "testing"
import "reflect"

type Case struct {
	Input []int
	Expected [][]int
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Input: []int{-1,0,1,2,-1,-4}, Expected: [][]int{{-1,0,1}, {-1,-1,2}}, },
	}

	for _, e := range cases {
		actual := threeSum(e.Input)

		if reflect.DeepEqual(actual, e.Expected) == false {
			t.Errorf("Given : %v, got: %v, actual: %v", e.Input, actual, e.Expected)
		}

	}
}
