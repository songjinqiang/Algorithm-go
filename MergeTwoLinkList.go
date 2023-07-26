package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个有序链表
func main() {
	linkList1 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}
	linkList2 := &ListNode{1, &ListNode{3, nil}}
	listResult := mergeTwoLists(linkList1, linkList2)
	printList(listResult)
}

func printList(list *ListNode) {
	for list != nil {
		println(list.Val)
		list = list.Next
	}
}

func mergeTwoLists(node1 *ListNode, node2 *ListNode) *ListNode {
	dummy := &ListNode{} // 虚拟节点 用于存储结果 因为current指针会不断后移
	current := dummy

	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			current.Next = node1
			node1 = node1.Next
		} else {
			current.Next = node2
			node2 = node2.Next
		}
		current = current.Next
	}
	if node1 != nil {
		current.Next = node1
	} else {
		current.Next = node2
	}

	return dummy.Next
}
