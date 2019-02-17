package leetcode

// 2.AddTwoNumbers
//
// Problem: https://leetcode.com/problems/add-two-numbers/
// Reference:
//
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	// merge into l1
	plusFlag := false
	for c1, c2 := l1, l2; ; {
		if c2 == nil {
			break
		}
		if plusFlag {
			c1.Val = c1.Val + c2.Val + 1
			plusFlag = false
		} else {
			c1.Val = c1.Val + c2.Val
		}
		if c1.Val >= 10 {
			c1.Val -= 10
			plusFlag = true
		}
		if c1.Next == nil && c2.Next != nil {
			c1.Next = new(ListNode)
		} else if c1.Next == nil && c2.Next == nil && plusFlag {
			c1.Next = new(ListNode)
			c1.Next.Val = 1
		} else if c1.Next != nil && c2.Next == nil && plusFlag {
			c2.Next = new(ListNode)
		}
		c1 = c1.Next
		c2 = c2.Next
	}

	return l1
}

// 3.LengthOfLongestSubstring
//
// Problem: https://leetcode.com/problems/longest-substring-without-repeating-characters/
//
func LengthOfLongestSubstring(s string) int {
	max := 0
	addIfNotExist := func(tm map[uint8]struct{}, c uint8) bool {
		_, ok := tm[c]
		tm[c] = struct{}{}
		return ok
	}

	for i := 0; i < len(s); i += 1 {
		cl := 0
		tmpMap := make(map[uint8]struct{})
		for j := i; j < len(s); j += 1 {
			if addIfNotExist(tmpMap, s[j]) {
				break
			}
			cl += 1
			if cl > max {
				max = cl
			}
		}
	}
	return max
}

// 5. LongestPalindrome
// Problem: https://leetcode.com/problems/longest-palindromic-substring/
func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var start, end int
	tryExpandToFull := func(str string, start int, mirror int) (int, int) {
		for start >= 0 && mirror < len(str) && str[start] == str[mirror] {
			start--
			mirror++
		}
		return start + 1, mirror - 1
	}
	for i, _ := range s[:len(s)-1] {
		if is, im := tryExpandToFull(s, i, i+1); (im - is) > (end - start) {
			start = is
			end = im
		}

		if is, im := tryExpandToFull(s, i, i); (im - is) > (end - start) {
			start = is
			end = im
		}

	}
	return s[start : end+1]
}
