package lintcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAplusB(t *testing.T) {
	var a int32 = -1
	var b int32 = 2
	var expect int32 = 1
	ret := aplusb(a, b)
	assert.Equal(t, expect, ret, "aplusb failed, input is %d and %d, expect %d but get %d", a, b, expect, ret)
}

func TestTrailingZeros(t *testing.T) {
	var input int64 = 105
	var expect int64 = 25
	ret := trailingZeros(input)
	assert.Equal(t, expect, ret, "trailingZeros failed, input is %d, expect %d, but get %d", input, expect, ret)
}

func TestMergeSortedArray(t *testing.T) {
	a := []int{7}
	b := []int{5, 7}
	expect := []int{5, 7, 7}
	ret := mergeSortedArray(a, b)
	assert.Equal(t, expect, ret, "mergeSortedArray failed, input is %v and %v, expect %v, but get %v", a, b, expect, ret)
}

func TestRotateString(t *testing.T) {
	raw := "abcdefg"
	var ret string
	cases := map[int]string{
		0: "abcdefg",
		1: "gabcdef",
		2: "fgabcde",
		3: "efgabcd",
	}
	for k, v := range cases {
		ret = rotateString(raw, k)
		assert.Equal(t, v, ret, "rotateString failed, input is %s and offset is %d, expect %s, but get %s", raw, k, v, ret)
	}
}

func TestRotateString2(t *testing.T) {
	raw := "abcdefg"
	var ret string
	cases := map[int]string{
		0: "abcdefg",
		1: "gabcdef",
		2: "fgabcde",
		3: "efgabcd",
	}
	for k, v := range cases {
		ret = rotateString2(raw, k)
		assert.Equal(t, v, ret, "rotateString2 failed, input is %s and offset is %d, expect %s, but get %s", raw, k, v, ret)
	}
}

func TestFizzBuzz(t *testing.T) {
	input := 15
	expect := []string{"1", "2", "fizz",
		"4", "buzz", "fizz",
		"7", "8", "fizz",
		"buzz", "11", "fizz",
		"13", "14", "fizz buzz",
	}
	ret := fizzBuzz(input)

	assert.Equal(t, expect, ret, "fizzBuzz failed, input is %d , expect %v, but get %v", input, expect, ret)

}

func TestStrStr(t *testing.T) {
	source := "tartarget"
	target := "target"
	expect := 3
	ret := strStr(source, target)
	assert.Equal(t, expect, ret, "strStr failed, input is %s, %s, expect %d, but get %d", source, target, expect, ret)
}

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4, 5, 10}
	target := 3
	expect := 2
	ret := binarySearch(nums, target)
	assert.Equal(t, expect, ret, "binarySearch failed, input is %v, target is %d, expect %d, but get %d", nums, target, expect, ret)
}
