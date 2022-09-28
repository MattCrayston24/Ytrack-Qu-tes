package main

import "fmt"

func main() {
	link := &List{}

	ListPushBack(link, 1)
	ListPushBack(link, 2)
	ListPushBack(link, 3)
	ListPushBack(link, 4)

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}


func ListPushBack(l *List, data interface{}) {
	var n *NodeL = &NodeL{Data: data}
	if l.Head == nil{
		l.Head = n
	}
	current := l.Head
	for current.Next != nil{
		current = current.Next
	}
	if n != l.Head{
		current.Next = n
	}
	l.Tail = n
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
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

func ListReverse(l *List) {
	j := l.Head
	if j == nil{
		return
	}
	var prev *NodeL
	l.Tail = l.Head
	for j != nil{
		next := j.Next
		j.Next = prev
		prev = j
		j = next
	}
	l.Head = prev
	return 
}