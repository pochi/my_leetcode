package leetcode015

import "sort"
import "strconv"
//import "fmt"

func threeSum(nums []int) [][]int {
	under_zeros := []int{}
	zeros := []int{}
	over_zeros := []int{}
	answers := [][]int{}
	nm := map[int]bool {}
	km := map[string]bool {}

	for _, i := range nums {
		if i < 0 {
			under_zeros = append(under_zeros, i)
		} else if i == 0 {
			zeros = append(zeros, i)
		} else {
			over_zeros = append(over_zeros, i)
		}

		nm[i] = true
	}

	var (
		c []int
		k string
	)

	if len(zeros) >= 3 {
		answers = append(answers, []int{0,0,0})
	}

	if len(zeros) >= 1 {
		for i:=0; i<len(under_zeros); i++ {
			if nm[under_zeros[i]*-1] {
				c = []int{under_zeros[i], 0, under_zeros[i]*-1}
				k = createKey(c)
				if km[k] == false {
					answers = append(answers, c)						
					km[k] = true
				}
			}
		}
	}


	for i:=0; i<(len(under_zeros)-1); i++ {
		for j:=i+1; j<len(under_zeros); j++ {
			if nm[(under_zeros[i] + under_zeros[j])*-1] {
				c = []int{under_zeros[i], under_zeros[j], (under_zeros[i] + under_zeros[j])*-1}
				sort.Ints(c)
				k = createKey(c)
				//fmt.Println(k)
				if km[k] == false {
					answers = append(answers, c)						
					km[k] = true
				}
			}
		}
	}

	for i:=0; i<(len(over_zeros)-1); i++ {
		for j:=i+1; j<len(over_zeros); j++ {
			if nm[(over_zeros[i] + over_zeros[j])*-1] {
				c = []int{over_zeros[i], over_zeros[j], (over_zeros[i] + over_zeros[j])*-1}
				sort.Ints(c)
				k = createKey(c)
				if km[k] == false {
					answers = append(answers, c)
					km[k] = true					
				}
				
			}
		}
	}

	return answers
}

func createKey(nums []int) string {
	return strconv.Itoa(nums[0]) + "_" + strconv.Itoa(nums[1]) + "_" + strconv.Itoa(nums[2])
}

// INFO: binary search tree. not used.
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