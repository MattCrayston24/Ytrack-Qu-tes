package main

import "fmt"

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}


func ListPushFront(l *List, data interface{}) {
	var n *NodeL = &NodeL{Data: data}
	if l.Head == nil{
		l.Head = n
	}else{
		n.Next = l.Head
		l.Head = n
	}
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	tab := 0
	if l.Head == nil{
		return 0
	}
	current := l.Head
	for current != nil{
		current = current.Next
		tab = tab+1
	}
	return tab
}