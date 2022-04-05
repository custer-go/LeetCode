package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	p := 0
	q := 0
	set := make(map[byte]bool, 0)
	maxLen := 0
	for q < n {
		c := s[q]
		if !set[c] {
			set[c] = true
			q++
			if q-p > maxLen {
				maxLen = q - p
			}
			continue
		}
		for set[c] {
			delete(set, s[p])
			p++
		}
	}
	return maxLen
}
