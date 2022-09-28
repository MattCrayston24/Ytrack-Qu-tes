package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}

func main() {
	link := &List{}

	ListPushBack(link, "I")
	ListPushBack(link, 1)
	ListPushBack(link, "something")
	ListPushBack(link, 2)

	fmt.Println("------list------")
	PrintList(link)
	ListClear(link)
	fmt.Println("------updated list------")
	PrintList(link)
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

func ListClear(l *List) {
	if l.Head != nil{
		l.Head = nil
	}
	if l.Tail != nil{
		l.Tail= nil
	}
}