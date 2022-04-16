package chain

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetListLength(head *ListNode) int {
	if head == nil {
		return 0
	}
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	return n
}
