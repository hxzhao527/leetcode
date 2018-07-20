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
