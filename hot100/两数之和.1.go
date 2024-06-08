package main

// 哈希
func towSum(nums []int, target int) []int {
	hashMap := make(map[int]int, 4)
	for i, v := range nums {
		if another, ok := hashMap[target-v]; ok {
			return []int{another, i}
		}
		hashMap[v] = i
	}
	return nil
}
