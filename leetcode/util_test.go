package leetcode

import (
	// "github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareListNode(t *testing.T) {
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
	//s2 := &ListNode{
	//	Val: 1,
	//	Next: & ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val:4,
	//			Next: nil,
	//		},
	//	},
	//}
	if ret := CompareListNode(s1, s1); !ret {
		t.Error("compare not well")
	}
}
