package main

import "slices"

// 1.排序
func longestConsecutiveBySort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)
	max, l := 1, 1
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if v == nums[i-1] {
			continue
		}
		if v == nums[i-1]+1 {
			l++
			max = func(max, l int) int {
				if max > l {
					return max
				}
				return l
			}(max, l)
		} else {
			l = 1
		}
	}
	return max
}
