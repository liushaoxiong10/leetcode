package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	if findMedianSortedArrays03(nums1, nums2) != 2 {
		fmt.Println(findMedianSortedArrays03(nums1, nums2))
	}
	nums3 := []int{1, 2}
	nums4 := []int{3, 4}
	if findMedianSortedArrays03(nums3, nums4) != 2.5 {
		fmt.Println(findMedianSortedArrays03(nums3, nums4))
	}
}

// 合并 两个有序列表
func findMedianSortedArrays01(nums1 []int, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	nums3 := make([]int, 0)

	if nums1Len == 0 {
		if (nums2Len % 2) == 0 {
			return float64(nums2[nums2Len/2-1]+nums2[nums2Len/2]) / 2
		}
		return float64(nums2[nums2Len/2])
	}

	if nums2Len == 0 {
		if (nums1Len % 2) == 0 {
			return float64(nums1[nums1Len/2-1]+nums1[nums1Len/2]) / 2
		}
		return float64(nums1[nums1Len/2])
	}

	var count, index1, index2 = 0, 0, 0
	for count != (nums1Len + nums2Len) {
		if index1 == nums1Len {
			for index2 != nums2Len {
				nums3 = append(nums3, nums2[index2])
				index2++
				count++
			}
			break
		}
		if index2 == nums2Len {
			for index1 != nums1Len {
				nums3 = append(nums3, nums1[index1])
				index1++
				count++
			}
			break
		}

		if nums1[index1] < nums2[index2] {
			nums3 = append(nums3, nums1[index1])
			index1++
		} else {
			nums3 = append(nums3, nums2[index2])
			index2++
		}
		count++
	}

	if (count % 2) == 0 {
		return float64(nums3[count/2-1]+nums3[count/2]) / 2
	}
	return float64(nums3[count/2])
}

// 伪合并
func findMedianSortedArrays02(nums1, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	countLen := nums1Len + nums2Len
	left, right, nums1Index, nums2Index := 0, 0, 0, 0
	for i := 0; i <= countLen/2; i++ {
		left = right
		if nums1Index < nums1Len && (nums2Index >= nums2Len || nums1[nums1Index] < nums2[nums2Index]) {
			right = nums1[nums1Index]
			nums1Index++
		} else {
			right = nums2[nums2Index]
			nums2Index++
		}
	}
	if (countLen % 2) == 0 {
		return float64(left+right) / 2
	}
	return float64(right)
}

// 二分法
func findMedianSortedArrays03(nums1, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	if nums1Len > nums2Len {
		return findMedianSortedArrays03(nums2, nums1)
	}
	iMin := 0
	iMax := nums1Len
	halfLen := (nums1Len + nums2Len + 1) / 2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxLeft := 0.0
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}
			if (nums1Len+nums2Len)%2 == 1 {
				return maxLeft
			}
			minRight := 0.0
			if i == nums1Len {
				minRight = float64(nums2[j])
			} else if j == nums2Len {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums1[i]), float64(nums2[j]))
			}
			return (maxLeft + minRight) / 2
		}
	}
	return 0
}

func findMedianSortedArrays04(nums1, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	left := (nums1Len + nums2Len + 1) / 2
	right := (nums1Len + nums2Len + 2) / 2
	return (getMedian(nums1, nums2, 0, nums1Len-1, 0, nums2Len-1, left) + getMedian(nums1, nums2, 0, nums1Len-1, 0, nums2Len-1, right)) / 2
}
func getMedian(nums1, nums2 []int, start1, end1, start2, end2, k int) float64 {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	if len1 > len2 {
		return getMedian(nums2, nums1, start2, end2, start1, end1, k)
	}
	if len1 == 0 {
		return float64(nums2[start2+k-1])
	}
	if k == 1 {
		return math.Min(float64(nums1[start1]), float64(nums2[start2]))
	}
	i := start1 + int(math.Min(float64(len1), float64(k/2))) - 1
	j := start2 + int(math.Min(float64(len2), float64(k/2))) - 1

	if nums1[i] > nums2[j] {
		return getMedian(nums1, nums2, start1, end1, j+1, end2, k-(j-start2+1))
	}
	return getMedian(nums1, nums2, i+1, end1, start2, end2, k-(i-start1+1))
}
