package leetcode

import "sort"

// 4. FindMedianSortedArrays
// Problem: https://leetcode.com/problems/median-of-two-sorted-arrays/
//
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	num := make([]int, 0, length)
	// how append implement? time is O(N) ?
	num = append(num, nums1...)
	num = append(num, nums2...)

	sort.Ints(num)
	if length%2 == 0 {
		return (float64(num[length/2-1]) + float64(num[length/2])) / float64(2.0)
	} else {
		return float64(num[length/2])
	}
}
