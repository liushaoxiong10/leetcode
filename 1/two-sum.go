package main

import "fmt"

// 暴力法
func twoSum(nums []int, target int) []int {
	for i, j := range nums {
		for o, k := range nums[i+1:] {
			if j+k == target {
				return []int{i, o + i + 1}
			}
		}
	}
	return nil
}

// map
func twoSumMap(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		result := target - nums[i]
		if index, ok := tmp[result]; ok {
			return []int{index, i}
		}
		tmp[nums[i]] = i

	}
	return nil
}

func main() {
	testExample := []int{0, 1, 2, 3}
	testResult := []int{1, 3}
	// 暴力法
	result := twoSum(testExample, 4)
	if intSliceEqualBCE(testResult, result) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}

	// map
	result = twoSumMap(testExample, 4)
	if intSliceEqualBCE(testResult, result) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}

}

func intSliceEqualBCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
