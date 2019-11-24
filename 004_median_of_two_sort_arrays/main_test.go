package leetcode004

import(
	"testing"
)

type Case struct {
	Nums1 []int
	Nums2 []int	
	Expected float64
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ Nums1: []int{1,3}, Nums2: []int{2}, Expected: 2.0, },
		Case{ Nums1: []int{1,2}, Nums2: []int{3,4}, Expected: 2.5, },		
	}

	for _, e := range cases {
		actual := findMedianSortedArrays(e.Nums1, e.Nums2)
		if actual != e.Expected {
			t.Errorf("got: %v, actual: %v", actual, e.Expected)
		}

	}
}
