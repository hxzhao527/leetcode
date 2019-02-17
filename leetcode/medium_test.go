package leetcode

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	s1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	s2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	expect := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}
	if ret := AddTwoNumbers(s1, s2); !CompareListNode(ret, expect) {
		t.Error("AddTwoNumbers failed")
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	data := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
	}
	for k, v := range data {
		if ret := LengthOfLongestSubstring(k); ret != v {
			t.Errorf("LengthOfLongestSubstring: %s, expect %d, got %d", k, v, ret)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	data := map[string]string{
		"babad": "bab",
		"cbbd":  "bb",
		"ccc":   "ccc",
		"cccc":  "cccc",
	}
	for k, v := range data {
		if ret := LongestPalindrome(k); ret != v {
			t.Errorf("LongestPalindrome %s, expect %s, got %s", k, v, ret)
		}
	}
}
