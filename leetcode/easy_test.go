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
