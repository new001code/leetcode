package main

func lengthOfLongestSubStrings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	ans := 0
	sBytes := []byte(s)
	hashSet := make(Set, 16)
	left := 0
	cur_len := 0
	for _, v := range sBytes {
		if hashSet.Contain(v) {
			for hashSet.Contain(v) {
				cur_len--
				hashSet.Delete(sBytes[left])
				left++
			}
		}
		hashSet.Add(v)
		cur_len++
		ans = max(ans, cur_len)
	}

	return ans
}

type Set map[byte]struct{}

func (s Set) Contain(key byte) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key byte) {
	s[key] = struct{}{}
}

func (s Set) Delete(key byte) {
	delete(s, key)
}
