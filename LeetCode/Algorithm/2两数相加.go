package algorithm

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p *ListNode = &ListNode{}
	var head *ListNode = p
	carry := 0
	for l1 != nil || l2 != nil {
		num1, num2 := 0, 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}
		num := num1 + num2 + carry
		num, carry = num%10, num/10
		p.Val = num
		if l1 != nil || l2 != nil {
			p.Next = &ListNode{}
			p = p.Next
		}
	}
	if carry > 0 {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = 1
	}
	return head
}

func Call_2() {

}
