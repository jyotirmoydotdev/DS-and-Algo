package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Printf("%d -> ", list.Val)
		list = list.Next
	}
	fmt.Println("nil")
}
