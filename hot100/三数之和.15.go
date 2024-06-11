package main

import "slices"

func threeNum(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for first := 0; first < n; first++ {
		if nums[first] > 0 {
			break
		}
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -nums[first]
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return ans
}
