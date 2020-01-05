package leetcode015

import "reflect"
import "sort"

func threeSum(nums []int) [][]int {
	under_zeros := []int{}
	zeros := []int{}
	over_zeros := []int{}
	answers := [][]int{}

	for _, i := range nums {
		if i < 0 {
			under_zeros = append(under_zeros, i)
		} else if i == 0 {
			zeros = append(zeros, i)
		} else {
			over_zeros = append(over_zeros, i)
		}
	}

	sort.Ints(under_zeros)
	sort.Ints(over_zeros)

	var d bool // d -> duplicate

	// Dirty program
	if len(zeros) >= 3 {
		answers = append(answers, []int{0,0,0})
	}

	var c []int
	
	if len(zeros) >= 1 {
		for i:=0; i<len(under_zeros); i++ {
			if find(over_zeros, under_zeros[i]*-1) {
				d = false
				c = []int{under_zeros[i], 0, under_zeros[i]*-1}
				for i:=0; i<len(answers); i++ {
					if reflect.DeepEqual(c, answers[i]) {
						d = true
					}
				}

				if d == false {
					answers = append(answers, c)						
				}
			}
		}
	}


	for i:=0; i<(len(under_zeros)-1); i++ {
		for j:=i+1; j<len(under_zeros); j++ {
			if find(over_zeros, (under_zeros[i] + under_zeros[j])*-1) {
				d = false
				c = []int{under_zeros[i], under_zeros[j], (under_zeros[i] + under_zeros[j])*-1}
				sort.Ints(c)
				for i:=0; i<len(answers); i++ {
					if reflect.DeepEqual(c, answers[i]) {
						d = true
					}
				}

				if d == false {
					answers = append(answers, c)					
				}
			}
		}
	}

	for i:=0; i<(len(over_zeros)-1); i++ {
		for j:=i+1; j<len(over_zeros); j++ {
			if find(under_zeros, (over_zeros[i] + over_zeros[j])*-1) {
				d = false
				c = []int{over_zeros[i], over_zeros[j], (over_zeros[i] + over_zeros[j])*-1}
				sort.Ints(c)
				for i:=0; i<len(answers); i++ {
					if reflect.DeepEqual(c, answers[i]) {
						d = true
					}
				}

				if d == false {
					answers = append(answers, c)					
				}
			}
		}
	}

	return answers
}

func find(nums []int, t int) bool {
	l := len(nums)
    s := 0
    e := l - 1
    var i int
    for {
        if e < s {
            return false
        }
        i = (s + e) / 2

        if nums[i] == t {
            return true
        }

        if nums[i] < t {
            s = i + 1
        } else {
            e = i - 1
        }
    }
}