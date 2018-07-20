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
	// control the order of output, because of hashMap
	com := func(a, b int) []int {
		if a < b {
			return []int{a, b}
		} else {
			return []int{b, a}
		}
	}
	for i, num := range nums {
		// num maybe negative, so the comparison is not need
		//if num <= target {
		if part[num] == nil {
			part[num] = make([]int, 0)
		}
		part[num] = append(part[num], i)
		//}
	}
	for num, in := range part {
		if in2, ok := part[target-num]; ok {
			// different num plus to target
			if len(in) == len(in2) && len(in) == 1 {
				if in[0] != in2[0] {
					return com(in[0], in2[0])
				}
			}
			// equivalent num plus to target
			// because of the description of question, there is only one combination,
			// so it's no need to slice the result.
			if len(in) > 1 {
				return in
			}
		}
	}
	return nil
}

// reverse
//
// Problem: https://leetcode-cn.com/problems/reverse-integer/description/
// Reference:
//
// @param abc: target to reverse
// @return: 0 or cba
//
func reverse(abc int) int {
	// can not compare abc to some value then return 0
	// e.g. 15000000000
	abs := func(a int) (int, bool) {
		if a < 0 {
			return -a, true
		}
		return a, false
	}

	x, flag := abs(abc)
	r := 0
	for ; x >= 10; x = x / 10 {
		r += x % 10
		r *= 10
	}
	if flag && -(x+r) > -(1<<31) {
		return -(x + r)
	}
	if !flag && (x+r) < 1<<31-1 {
		return x + r
	}
	return 0
}

// isPalindrome
//
// Problem: https://leetcode-cn.com/problems/palindrome-number/description/
// Reference:
//
// @param abc: target to reverse
// @return:
//
func isPalindrome(abc int) bool {
	if abc >= 0 && abc <= 9 {
		return true
	}
	if abc < 0 || abc%10 == 0 {
		return false
	}
	x := abc
	r := 0
	for ; x >= 10; x = x / 10 {
		r += x % 10
		r *= 10
	}
	return abc == x+r
}
