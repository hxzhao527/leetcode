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

// romanToInt
//
// Problem: https://leetcode-cn.com/problems/roman-to-integer/description/
// Reference:
//
// @param s: roman number
// @return:
//
func romanToInt(s string) int {
	dataSet := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	ret := 0
	for i := len(s) - 1; i > 0; i-- {
		if dataSet[string(s[i])] > dataSet[string(s[i-1])] {
			ret += dataSet[string(s[i])] - dataSet[string(s[i-1])]
			if i == 1 {
				return ret
			}
			i--
			continue
		} else {
			ret += dataSet[string(s[i])]
		}
	}
	ret += dataSet[string(s[0])]
	return ret
}

// longestCommonPrefix
//
// Problem: https://leetcode-cn.com/problems/longest-common-prefix/description/
// Reference:
//
// @param s: roman number
// @return:
//
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	prefix := make([]uint8, 0)

OutLoop:
	for i := 0; i < len(strs[0]); i++ {
		temp := strs[0][i]
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) <= i {
				break OutLoop
			}
			if strs[j][i] == temp {
				continue
			} else {
				break OutLoop
			}
		}
		prefix = append(prefix, temp)
	}
	return string(prefix)
}

// isValid
//
// Problem: https://leetcode-cn.com/problems/valid-parentheses/description/
// Reference:
//
// @param s: roman number
// @return:
//
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	//if len(s) == 1{
	//	return false
	//}
	//length := len(s)
	stack := make([]byte, 0)

	for v := range s {
		if s[v] == '(' || s[v] == '{' || s[v] == '[' {
			stack = append(stack, s[v])
		} else {
			if len(stack) == 0 {
				return false
			}
			switch s[v] {
			case ')':
				if stack[len(stack)-1] != '(' {
					return false
				}
				break
			case '}':
				if stack[len(stack)-1] != '{' {
					return false
				}
				break
			case ']':
				if stack[len(stack)-1] != '[' {
					return false
				}
				break
			default:
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
