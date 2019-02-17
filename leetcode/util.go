package leetcode

func CompareListNode(l1 *ListNode, l2 *ListNode) bool {

	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		if l1.Val != l2.Val {
			return false
		}
	}

	if l1 == nil && l2 != nil {
		return false
	}
	if l1 != nil && l2 == nil {
		return false
	}

	return true
}
