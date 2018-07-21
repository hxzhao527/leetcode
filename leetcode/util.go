package leetcode

func CompareListNode(l1 *ListNode, l2 *ListNode) int {

	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		if l1.Val > l2.Val {
			return 1
		} else if l1.Val == l2.Val {
			continue
		} else {
			return -1
		}
	}

	if l1 == nil && l2 != nil {
		return -1
	}
	if l1 != nil && l2 == nil {
		return 1
	}

	return 0
}
