package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{3, 2, 4}
	target := 6
	expect := []int{1, 2}
	ret := twoSum(input, target)
	assert.Equal(t, expect, ret, "twoSum failed")
}

func TestReverse(t *testing.T) {
	input := 1534236469
	expect := 0
	ret := reverse(input)
	assert.Equal(t, expect, ret, "reverse failed")
}

func TestIsPalindrome(t *testing.T) {
	dataSet := map[int]bool{
		121:  true,
		-121: false,
		10:   false,
	}
	for in, flag := range dataSet {
		assert.Equal(t, flag, isPalindrome(in), "isPalindrome failed")
	}
}

func TestRomanToInt(t *testing.T) {
	dataSet := map[string]int{
		"III":     3,
		"IV":      4,
		"IX":      9,
		"LVIII":   58,
		"MCMXCIV": 1994,
	}
	for in, val := range dataSet {
		assert.Equal(t, val, romanToInt(in), "romanToInt failed")
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	dataSet := map[string][]string{
		"fl": {"flower", "flow", "flight"},
		//"": {"dog","racecar","car"},
		"a": {"aa", "a"},
		"":  nil,
	}
	for pre, set := range dataSet {
		assert.Equal(t, pre, longestCommonPrefix(set), "longestCommonPrefix failed")
	}
}

func TestIsValid(t *testing.T) {
	dataSet := map[string]bool{
		"(":      false,
		"()":     true,
		"()[]{}": true,
		"(]":     false,
		"([)]":   false,
		"{[]}":   true,
	}
	for in, val := range dataSet {
		assert.Equal(t, val, isValid(in), "isValid failed, input is %s", in)
	}
}

func TestMergeTwoLists(t *testing.T) {
	s1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	s2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	expect := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}
	ret := mergeTwoLists(s1, s2)
	assert.Equal(t, 0, CompareListNode(expect, ret), "merge failed")

}

func TestRemoveDuplicates(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expect := []int{0, 1, 2, 3, 4}
	length := 5
	assert.Equal(t, length, removeDuplicates(input), "remove failed")
	assert.Equal(t, expect, input[:length], "remove result not equal")
}

func TestRemoveElement(t *testing.T) {
	input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	expect := []int{0, 1, 3, 0, 4}
	length := 5
	assert.Equal(t, length, removeElement(input, val), "remove failed")
	assert.Equal(t, expect, input[:length], "remove result not equal")
}

func TestStrStr(t *testing.T) {
	source := "tartarget"
	target := "target"
	expect := 3
	ret := strStr(source, target)
	assert.Equal(t, expect, ret, "strStr failed, input is %s, %s, expect %d, but get %d", source, target, expect, ret)
}

func TestSearchInsert(t *testing.T) {
	input := []int{1, 3}
	target := 3
	expect := 1
	assert.Equal(t, expect, searchInsert(input, target), "searchInsert failed")
}

func TestCountAndSay(t *testing.T) {
	countAndSay(1)
	t.Log("cann't understand the question!")
}

func TestMaxSubArray(t *testing.T) {
	input := []int{-1, -2, 1}
	expect := 1
	assert.Equal(t, expect, maxSubArray(input), "maxSubArray failed")
}
func TestLengthOfLastWord(t *testing.T) {
	lengthOfLastWord("")
	t.Log("cann't understand the question!")
}
