package leetcode016

import "math"
import "sort"

func threeSumClosest(nums []int, t int) int {
	s := make(map[int]bool)

	for i:=0; i<len(nums)-2; i++ {
		for j:=i+1; j<len(nums)-1; j++ {
			for h:=j+1; h<len(nums); h++ {
				s[nums[i]+nums[j]+nums[h]] = true
			}
		}
	}


	sums := sortKeys(s)
	return findClosest(sums, t)
}

func sortKeys(s map[int]bool) []int {
	sums := []int{}
	for k, _ := range s {
		sums = append(sums, k)
	}
	sort.Ints(sums)
	return sums
}

func findClosest(sums []int, t int) int {
	var (
		d int
		a int
	)
	c := 10000
	for _, s := range sums {
		d = int(math.Abs((float64(s) - float64(t))))
		if c > d {
			c = d
			a = s
		}
	}
	return a
}