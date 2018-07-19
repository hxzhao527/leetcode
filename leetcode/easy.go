/*
Package leetcode solve leetcode problems with golang

solved by hxzhao(haoxiangzhao@outlook.com)

*/
package leetcode

// twoSum
//
// Problem: https://leetcode-cn.com/problems/two-sum/description/
// Reference:
//
// @param nums: An integer array
// @param target: An integer
// @return: The position of a and b
//
func twoSum(nums []int, target int) []int {
	part := make(map[int][]int)
	//
	com := func(a, b int) []int {
		if a < b {
			return []int{a, b}
		} else {
			return []int{b, a}
		}
	}
	for i, num := range nums {
		// maybe -3 and +3
		//if num <= target {
		if part[num] == nil {
			part[num] = make([]int, 0)
		}
		part[num] = append(part[num], i)
		//}
	}
	for num, in := range part {
		if in2, ok := part[target-num]; ok {
			if len(in) == len(in2) && len(in) == 1 {
				if in[0] != in2[0] {
					return com(in[0], in2[0])
				}
			}
			if len(in) > 1 {
				return in
			}
		}
	}
	return nil
}
