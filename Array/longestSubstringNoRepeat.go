package main

func lengthOfLongestSubstring(s string) int {
	count := 0
	if len(s) < 1 {
		return count
	}
	sb := []byte(s)
	start := 0
	for i := 0; i < len(sb); i++ {
		for j := start; j < i; j++ {
			if sb[j] == sb[i] {
				start = j + 1
				break

			}
		}
		count = max(count, i-start+1)
	}

	return count
}
