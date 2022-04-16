package chain

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	link := &ListNode{0, head}
	len := GetListLength(head) //注意：二者指向同一块内存地址，但是二者不是引用关系
	fmt.Println(head.Val)
	inter := link
	removeNodeIndex := len - n + 1
	index := 1
	for inter.Next != nil && index < removeNodeIndex {
		inter = inter.Next
		index++
	}
	inter.Next = inter.Next.Next
	return link.Next
}
