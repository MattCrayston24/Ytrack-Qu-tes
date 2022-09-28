package main

import "fmt"

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "hello1")
	ListPushBack(link, "hello2")
	ListPushBack(link, "hello3")

	found := ListFind(link, interface{}("hello2"), CompStr)

	fmt.Println(found)
	fmt.Println(*found)
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

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head
	if current == nil{
		return nil
	}
	for current != nil{
		if comp(current.Data, ref) == true{
			return &current.Data
		}
		current = current.Next
	}
	return nil
}