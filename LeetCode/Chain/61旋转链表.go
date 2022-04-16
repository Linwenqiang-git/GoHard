package chain

import "fmt"

func rotateRight(head *ListNode, k int) *ListNode {
	n := GetListLength(head)
	if n == 0 {
		return head
	}
	k = k % n //去掉周期性
	if k == 0 {
		return head
	}
	newNum := n - k
	startNum := 1
	oldHead := head
	for startNum < newNum && head != nil {
		startNum++
		head = head.Next
	}
	//将列表一拆为二
	newHead := head.Next
	fmt.Println(newHead.Val)
	head.Next = nil
	p := newHead
	for p.Next != nil {
		p = p.Next
	}
	p.Next = oldHead
	return newHead
}
