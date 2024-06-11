package main

// 最简单的思路。 O(N^2)
// 会超时
func maxArea(height []int) int {
	max := 0
	for i, v := range height {
		for j := i + 1; j < len(height); j++ {
			max = getMax(max, getMin(v, height[j])*(j-i))
		}
	}
	return max
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxAreaPlus(height []int) int {
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		max = getMax(max, getMin(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
