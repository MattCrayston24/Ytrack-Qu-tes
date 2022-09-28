package main

import "fmt"

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "a")
	ListPushBack(link, "b")
	ListPushBack(link, "c")
	ListPushBack(link, "d")
	fmt.Println("-----first List------")
	PrintList(link)

	ListPushBack(link2, "e")
	ListPushBack(link2, "f")
	ListPushBack(link2, "g")
	ListPushBack(link2, "h")
	fmt.Println("-----second List------")
	PrintList(link2)

	fmt.Println("-----Merged List-----")
	ListMerge(link, link2)
	PrintList(link)
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

func ListMerge(l1 *List, l2 *List) {
	current := l1.Head
	current2 := l2.Head
	for current != nil{
		if current.Next != nil{
			l1.Tail.Next = current2
		}
		return
	}
}

