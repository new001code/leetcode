package main

import (
	"slices"
)

func sortStringChar(str string) string {
	// 将字符串转为byte切片
	tempByteSlice := []byte(str)
	// 排序
	slices.Sort(tempByteSlice)
	// 返回
	return string(tempByteSlice)
}

// 哈希 + 排序
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	hashMap := make(map[string][]string, 16)
	for _, str := range strs {
		sortStr := sortStringChar(str)
		if slice, ok := hashMap[sortStr]; ok {
			slice = append(slice, str)
			hashMap[sortStr] = slice
		} else {
			hashMap[sortStr] = []string{str}
		}
	}

	res := make([][]string, len(hashMap))
	i := 0
	for _, v := range hashMap {
		res[i] = v
		i++
	}
	return res
}
