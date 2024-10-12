package main

import (
	"fmt"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	// Edge case if both list does not have any element
	// return nil
	if list1 == nil && list2 == nil {
		return nil
	}

	// If the list1 is nil than return list2
	// Because one of the list has some element
	if list1 == nil {
		return list2
	}

	// If the list 2 is nil return the list1
	// Because one of the list has some element
	if list2 == nil {
		return list1
	}

	// So we are making the reference pointer
	ptr1, ptr2 := list1, list2

	// allocate memory for a new ListNode
	result := &ListNode{}

	temp := result

	for ptr1 != nil && ptr2 != nil {
		printList(result)
		printList(temp)
		if ptr1.Val < ptr2.Val {
			temp.Next = ptr1
			temp = temp.Next
			ptr1 = ptr1.Next
		} else {
			temp.Next = ptr2
			temp = temp.Next
			ptr2 = ptr2.Next
		}
	}

	for ptr1 != nil {
		temp.Next = ptr1
		temp = temp.Next
		ptr1 = ptr1.Next
	}

	for ptr2 != nil {
		temp.Next = ptr2
		temp = temp.Next
		ptr2 = ptr2.Next
	}

	result = result.Next
	return result
}

func RunMergeTwoLists() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}
	list2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: nil}}}

	fmt.Println("List 1:")
	printList(list1)
	fmt.Println("List 2:")
	printList(list2)

	mergedList := mergeTwoLists(list1, list2)
	fmt.Println("Merged List:")
	printList(mergedList)
}
