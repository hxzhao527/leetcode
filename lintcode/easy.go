/*
Package lintcode solve lintcode problems with golang

solved by hxzhao(haoxiangzhao@outlook.com)

*/
package lintcode

import (
	"bytes"
	"fmt"
)

// aplusb
//
// Problem: https://www.lintcode.com/problem/a-b-problem/description
// Reference: https://en.wikipedia.org/wiki/Bitwise_operation#Bitwise_operators
//
// @param a: An integer
// @param b: An integer
// @return: The sum of a and b
//
func aplusb(a int32, b int32) int32 {
	result := a ^ b       // + without carry 0+0=0, 0+1=1+0=1, 1+1=0
	carry := (a & b) << 1 // 1+1=2
	if carry != 0 {
		return aplusb(result, carry)
	}
	return result
}

// trailingZeros
//
// Problem: https://www.lintcode.com/problem/trailing-zeros/description
// Reference:
//
// @param n: A long integer
// @return: An integer, denote the number of trailing zeros in n!
//
func trailingZeros(n int64) int64 {
	result := n / 5
	if result%5 == 0 || result > 5 {
		return result + trailingZeros(result)
	}
	return result
}

// mergeSortedArray
//
// Problem: https://www.lintcode.com/problem/merge-two-sorted-arrays/description
// Reference:
//
// @param A: sorted integer array A
// @param B: sorted integer array B
// @return: A new sorted integer array
//
func mergeSortedArray(A []int, B []int) []int {
	aLength, bLength := len(A), len(B)

	if aLength == 0 {
		return B
	}
	if bLength == 0 {
		return A
	}

	aMin, aMax := A[0], A[aLength-1]
	bMin, bMax := B[0], B[bLength-1]

	if aMax <= bMin {
		A = append(A, B...)
		return A
	}
	if bMax <= aMin {
		B = append(B, A...)
		return B
	}
	result := make([]int, aLength+bLength, aLength+bLength)
	i, j, k := 0, 0, 0

	for ; i < aLength && j < bLength; k++ {
		if A[i] < B[j] {
			result[k] = A[i]
			i++
		} else {
			result[k] = B[j]
			j++
		}
	}
	if i < aLength {
		copy(result[k:], A[i:])
	}
	if j < bLength {
		copy(result[k:], B[j:])
	}
	return result
}

// rotateString
//
// Problem: https://www.lintcode.com/problem/rotate-string/description
// Reference:
//
// @param str: raw string
// @param offset: An integer
// @return: nothing
//
func rotateString(raw string, offset int) string {
	var buffer bytes.Buffer
	length := len(raw)
	if length == 0 {
		return ""
	}
	if offset > length {
		offset = offset % length
	}
	buffer.WriteString(raw[length-offset:])
	buffer.WriteString(raw[:length-offset])
	return buffer.String()
}

// rotateString
//
// Problem: https://www.lintcode.com/problem/rotate-string/description
// Reference:
//
// Rotate in-place with O(1) extra memory.
//
// @param str: raw string
// @param offset: An integer
// @return: nothing
//
func rotateString2(raw string, offset int) string {

	revert := func(raw []rune) {
		length := len(raw)
		for i := 0; i < length/2; i++ {
			raw[i], raw[length-1-i] = raw[length-1-i], raw[i]
		}
	}

	length := len(raw)
	if length == 0 {
		return ""
	}
	if offset > length {
		offset = offset % length
	}
	tmp := []rune(raw)
	revert(tmp)
	revert(tmp[:offset])
	revert(tmp[offset:])
	return string(tmp)
}

// fizzBuzz
//
// Problem: https://www.lintcode.com/problem/fizz-buzz/description
// Reference:
//
// do it with only one if statement
//
// @param n: An integer
// @return: A list of strings.
//
func fizzBuzz(n int) []string {
	result := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result = append(result, "fizz buzz")
		case i%3 == 0:
			result = append(result, "fizz")
		case i%5 == 0:
			result = append(result, "buzz")
		default:
			result = append(result, fmt.Sprintf("%d", i))
		}
	}
	return result
}

// strStr
//
// Problem: https://www.lintcode.com/problem/implement-strstr/description
// Reference: http://www.ruanyifeng.com/blog/2013/05/Knuth–Morris–Pratt_algorithm.html
// Reference: https://www.cnblogs.com/yjiyjige/p/3263858.html
//
// @param source: source string to be scanned.
// @param target: target string containing the sequence of characters to match
// @return a index to the first occurrence of target in source, or -1  if target is not part of source.
//kmp

func strStr(source, target string) int {
	if len(target) == 0 {
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
	}(target)
	//fmt.Printf("next map: %v", nextMap)

	for i, j := 0, 0; i < len(source); {
		if source[i] == target[j] {
			i++
			j++
			if j < len(target) {
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

// binarySearch
//
// Problem: https://www.lintcode.com/problem/first-position-of-target/description
// Reference:
//
// @param nums: The integer array.
// @param target: Target to find.
// @return: The first position of target. Position starts from 0.
//
func binarySearch(nums []int, target int) int {
	//index := -1
	minIndex := 0
	maxIndex := len(nums) - 1
	if nums[minIndex] > target || nums[maxIndex] < target {
		return -1
	}
	avgIndex := len(nums) / 2
	for minIndex != avgIndex {
		if nums[maxIndex] != target {
			if target > nums[avgIndex] {
				minIndex = avgIndex
			} else {
				maxIndex = avgIndex
			}
			avgIndex = (maxIndex + minIndex + 1) / 2
		} else {
			return maxIndex
		}
	}
	return -1
}
