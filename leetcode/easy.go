/*
Package leetcode solve leetcode problems with golang

solved by hxzhao(haoxiangzhao@outlook.com)

*/
package leetcode

import "sort"

// 1.TwoSum
//
// Problem: https://leetcode-cn.com/problems/two-sum/description/
// Reference:
//
func TwoSum(nums []int, target int) []int {
	part := make(map[int][]int)
	for index, num := range nums {
		if _, ok := part[num]; !ok {
			part[num] = make([]int, 0)
		}
		part[num] = append(part[num], index)
	}

	for num, index := range part {
		if _, ok := part[target-num]; ok {
			if num == target-num {
				if len(index) > 1 {
					return index
				} else {
					continue
				}
			} else {
				ret := make([]int, 0)
				// almost one solution
				ret = append(ret, index...)
				ret = append(ret, part[target-num]...)
				sort.Ints(ret)
				return ret
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

// mergeTwoLists
//
// Problem: https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
// Reference:
//
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var start1 = new(ListNode)
	position := start1

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			position.Next = l2
			l2 = l2.Next
		} else if l1.Val == l2.Val {
			position.Next = l1
			l1 = l1.Next
		} else {
			position.Next = l1
			l1 = l1.Next
		}
		position = position.Next
	}
	if l1 != nil {
		position.Next = l1
	}
	if l2 != nil {
		position.Next = l2
	}

	return start1.Next
}

// removeDuplicates
//
// Problem: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/description/
// Reference:
//
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	curVal := nums[0]
	curInd := 0
	for i := range nums {
		if curVal == nums[i] {
			continue
		} else {
			nums[curInd] = curVal

			curInd++
			curVal = nums[i]
		}
	}
	nums[curInd] = curVal
	return curInd + 1
}

// removeElement
//
// Problem: https://leetcode-cn.com/problems/remove-element/description/
// Reference:
//
func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	curInd := 0
	for i := range nums {
		if nums[i] == val {
			continue
		} else {
			nums[curInd] = nums[i]
			curInd++
		}
	}
	return curInd
}

// strStr
//
// Problem: https://leetcode-cn.com/problems/implement-strstr/description/
// Reference: http://www.ruanyifeng.com/blog/2013/05/Knuth–Morris–Pratt_algorithm.html
// Reference: https://www.cnblogs.com/yjiyjige/p/3263858.html
//
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	nextMap := func(raw string) []int {
		nextMap := make([]int, len(raw))
		nextMap[0] = -1
		k := -1
		for j := 0; j < len(raw)-1; {
			if k == -1 || raw[j] == raw[k] {
				k++
				j++
				nextMap[j] = k
			} else {
				k = nextMap[k]
			}
		}
		//nextMap[0] = 0
		return nextMap
	}(needle)
	//fmt.Printf("next map: %v", nextMap)

	for i, j := 0, 0; i < len(haystack); {
		if haystack[i] == needle[j] {
			i++
			j++
			if j < len(needle) {
				continue
			} else {
				return i - j
			}
		} else {
			next := nextMap[j]
			if j == 0 {
				i = i + (j - next)
				j = 0
			} else {
				j = next
			}
		}
	}
	return -1
}

// searchInsert
//
// Problem: https://leetcode-cn.com/problems/search-insert-position/description/
// Reference:
//
func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	offset := 0
	for mid := len(nums) / 2; len(nums) > 2; mid = len(nums) / 2 {
		if target > nums[mid] {
			offset = offset + mid
			nums = nums[mid:]
		} else {
			nums = nums[:mid+1]
		}
	}
	if nums[0] == target {
		return offset
	} else {
		return offset + 1
	}
}

// searchInsert
//
// Problem: https://leetcode-cn.com/problems/count-and-say/description/
// Reference:
//
func countAndSay(n int) string {
	n = n
	return ""
}

// searchInsert
//
// Problem: https://leetcode-cn.com/problems/maximum-subarray/description/
// Reference: http://www.kitabxana.net/files/books/file/1354098277.pdf
//
func maxSubArray(nums []int) int {

	maxSubSum, thisSubSum := 0, 0
	// because int var default value is 0,
	// so, if all elements in slice is negative, the max value should be the first, not 0
	start := true
	for _, num := range nums {
		thisSubSum += num
		if thisSubSum > maxSubSum || start {
			maxSubSum = thisSubSum
		}
		if thisSubSum < 0 {
			thisSubSum = 0
		}
		start = false
	}
	return maxSubSum
}

// searchInsert
//
// Problem: https://leetcode-cn.com/problems/length-of-last-word/description/
// Reference:
//
func lengthOfLastWord(s string) int {
	s = s
	return 0
}

// searchInsert
//
// Problem: https://leetcode-cn.com/problems/plus-one/description/
// Reference:
//
func plusOne(digits []int) []int {
	length := len(digits)
	if length == 0 {
		return []int{1}
	}
	if digits[length-1] != 9 {
		digits[length-1] += 1
		return digits
	}

	digits[len(digits)-1] = 0

	if len(digits) != 1 {
		plusOne(digits[:length-1])
		//digits = append(digits, 0)
	} else {
		digits = append(digits, 0)
		copy(digits[1:], digits[:])
		digits[0] = 1
	}
	// make up the tail 0
	if length >= 2 && digits[0] == 1 && digits[1] == 0 {
		digits = append(digits, 0)
	}
	return digits
}
