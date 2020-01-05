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
			for j:=0; j<len(over_zeros); j++ {
				if (under_zeros[i] + over_zeros[j]) == 0 {
					d = false
					c = []int{under_zeros[i], 0, over_zeros[j]}
					for i:=0; i<len(answers); i++ {
						if reflect.DeepEqual(c, answers[i]) {
							d = true
						}
					}

					if d == false {
						answers = append(answers, c)						
					}
					break
				}

				if (under_zeros[i]*-1) < over_zeros[j] {
					break
				}
			}
		}
	}


	for i:=0; i<(len(under_zeros)-1); i++ {
		for j:=i+1; j<len(under_zeros); j++ {
			for h:=0; h<len(over_zeros); h++ {
				if under_zeros[i] + under_zeros[j] + over_zeros[h] == 0 {
					d = false
					c = []int{under_zeros[i], under_zeros[j], over_zeros[h]}
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

				if ((under_zeros[i]+under_zeros[j])*-1) < over_zeros[h] {
					break
				}
			}
		}
	}

	for i:=0; i<(len(over_zeros)-1); i++ {
		for j:=i+1; j<len(over_zeros); j++ {
			for h:=0; h<len(under_zeros); h++ {
				if over_zeros[i] + over_zeros[j] + under_zeros[h] == 0 {
					d = false
					c = []int{over_zeros[i], over_zeros[j], under_zeros[h]}
					sort.Ints(c)
					for i:=0; i<len(answers); i++ {
						if reflect.DeepEqual(c, answers[i]) {
							d = true
						}
					}

					if d == false {
						answers = append(answers, c)					
					}
					break
				}

				if (over_zeros[i]+over_zeros[j]) > (under_zeros[h]*-1) {
					break
				}
			}
		}
	}

	return answers
}