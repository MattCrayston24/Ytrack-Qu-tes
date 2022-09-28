package main

import "fmt"

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *NodeI

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(ListSort(link))
}


type NodeI struct {
	Data int
	Next *NodeI
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
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

func ListSort(l *NodeI) *NodeI {
	current := l
	if l == nil{
		return nil
	}
	var prev *NodeI
	for current != nil{
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}