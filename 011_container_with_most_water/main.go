package leetcode011

//import "fmt"

func maxArea(height []int) int {
	max, w := 0, 0

	for i, h := range height {
		w = calc(h, height[i+1:])
		if max < w {
			max = w
		}
	}

	return max
}

func calc(h int, height []int) int {
	m, ch ,cm := 0, 0, 0
	for i, h2 := range height {
		if h > h2 {
			ch = h2
		} else {
			ch = h
		}
		
		cm = ch * (i+1)
		if m < cm {
			m = cm
		}
	}

	return m
}

