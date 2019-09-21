package main

import "fmt"

func main() {
	a := `abcda`

	fmt.Println(lengthOfLongestSubstring02(a))

}

func lengthOfLongestSubstring01(s string) int {
	maxLength := 0
	for i := 0; i < len(s); i++ {
		hash := make(map[byte]interface{})
		for j := i; j < len(s); j++ {
			if _, ok := hash[s[j]]; ok {
				break
			}
			hash[s[j]] = struct{}{}
		}
		if maxLength < len(hash) {
			maxLength = len(hash)
		}
	}
	if len(s) > 0 && maxLength == 0 {
		maxLength = 1
	}
	return maxLength
}

func lengthOfLongestSubstring02(s string) int {
	maxLength := 0
	hash := make(map[byte]int)
	i := 0
	for j := 0; j < len(s); j++ {
		if index, ok := hash[s[j]]; ok && index >= i {
			i = index + 1
		} else {
			newLength := j - i + 1
			if maxLength < newLength {
				maxLength = newLength
			}
		}
		hash[s[j]] = j
	}
	if len(s) > 0 && maxLength == 0 {
		maxLength = 1
	}
	return maxLength
}
