package chain

import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	data := []int{}
	newList := &ListNode{}
	for l1 != nil || l2 != nil {
		//有一个链表是空的，那么后面的直接补上
		if l1 == nil || l2 == nil {
			for l1 != nil {
				data = append(data, l1.Val)
				l1 = l1.Next
			}
			for l2 != nil {
				data = append(data, l2.Val)
				l2 = l2.Next
			}
			break
		}
		if l1.Val <= l2.Val {
			data = append(data, l1.Val)
			//newList = l1
			l1 = l1.Next
		} else if l1.Val > l2.Val {
			data = append(data, l2.Val)
			//newList = l2
			l2 = l2.Next
		}

		//newList = newList.Next
	}
	fmt.Printf("%v \n", data)
	var firstPoint *ListNode
	for index, value := range data {

		// if newList == nil {
		// 	newList = &ListNode{}
		// }

		newList.Val = value
		newList.Next = &ListNode{}
		if index == 0 {
			firstPoint = newList
		}

		if index+1 == len(data) {
			newList.Next = nil
			break
		}
		newList = newList.Next
	}
	return firstPoint
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	newList := &ListNode{}
	var firstPoint *ListNode
	index := 0
	for l1 != nil || l2 != nil {

		//有一个链表是空的，那么后面的直接补上
		if l1 == nil || l2 == nil {
			for l1 != nil {
				newList.Val = l1.Val
				if index == 0 {
					firstPoint = newList
				}
				index++
				l1 = l1.Next
				if l1 != nil {
					newList.Next = &ListNode{}
					newList = newList.Next
				}

			}
			for l2 != nil {
				newList.Val = l2.Val
				if index == 0 {
					firstPoint = newList
				}
				index++
				l2 = l2.Next
				if l2 != nil {
					newList.Next = &ListNode{}
					newList = newList.Next
				}

			}
			break
		}
		newList.Next = &ListNode{}
		if l1.Val <= l2.Val {
			newList.Val = l1.Val
			l1 = l1.Next
		} else if l1.Val > l2.Val {
			newList.Val = l2.Val
			l2 = l2.Next
		}
		if index == 0 {
			firstPoint = newList
		}

		newList = newList.Next
		index++
	}
	return firstPoint
}

func Call_21() {
	l1_3 := &ListNode{Val: 5, Next: nil}
	l1_2 := &ListNode{Val: 2, Next: l1_3}
	l1_1 := &ListNode{Val: 1, Next: l1_2}

	l2_3 := &ListNode{Val: 4, Next: nil}
	l2_2 := &ListNode{Val: 3, Next: l2_3}
	l2_1 := &ListNode{Val: 1, Next: l2_2}

	data := mergeTwoLists2(l1_1, l2_1)
	fmt.Println("输出结果：")
	for data != nil {
		fmt.Printf("%d ", data.Val)
		data = data.Next
	}
}
