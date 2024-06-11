package main

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		// 找到下一个0
		if nums[i] != 0 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			// 找到下一个非零
			if nums[j] == 0 {
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
			break
		}

	}
}
