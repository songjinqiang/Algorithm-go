package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

func Insert(data interface{}, position *Node) {
	// 新建一个节点
	tmpCell := new(Node)
	if tmpCell == nil {
		fmt.Println("err: out of space")
	}
	// 给新建的节点的值域赋值为传入的data值
	tmpCell.Data = data
	// 新建的节点的next指针指向的是指定节点position的next
	tmpCell.Next = position.Next
	// 指定节点position的后继节点变成了新建的节点
	position.Next = tmpCell
}

// 查找指定节点的前一个节点
func FindPrevious(data interface{}, node *Node) *Node {
	tmp := node
	for tmp.Next != nil && tmp.Next.Data != data {
		tmp = tmp.Next
	}
	return tmp
}

// 查找某个值在哪个节点
func Find(data interface{}, node *Node) *Node {
	tmp := node
	for tmp.Data != data {
		tmp = tmp.Next
	}
	return tmp
}

func Delete(data interface{}, node *Node) {
	preview := FindPrevious(data, node)
	tmp := Find(data, node)
	preview.Next = tmp.Next
	tmp = nil
}

func PrintList(list *Node) {
	p := list
	for p != nil {
		fmt.Printf("%v-%p\n", p.Data, p.Next)
		p = p.Next
	}
}

func main() {
	headNode := &Node{
		Data: 0,
		Next: nil,
	}
	list := headNode
	cur := headNode
	for i := 11; i < 20; i++ {
		Insert(i, cur)
		cur = cur.Next
	}
	Delete(13, list)
	PrintList(list)
}
