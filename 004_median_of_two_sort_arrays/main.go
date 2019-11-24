package leetcode004

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergeArray := append(nums1, nums2...)
	sort.Ints(mergeArray)
	medianNumber := len(mergeArray) / 2

	if len(mergeArray) % 2 == 1 {
		return float64(mergeArray[medianNumber])
	}
	
	return float64(mergeArray[medianNumber - 1] + mergeArray[medianNumber]) / 2
}
