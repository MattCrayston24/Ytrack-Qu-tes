package main

import (
    "fmt"
)

func PrintList(l *NodeI) {
    it := l
    for it != nil{
        fmt.Print(it.Data, " -> ")
        it = it.Next
    }
    fmt.Print(nil, "\n")
}
type NodeI struct {
    Data int
    Next *NodeI
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

    link = listPushBack(link, 1)
    link = listPushBack(link, 4)
    link = listPushBack(link, 9)

    PrintList(link)

    link = SortListInsert(link, -2)
    link = SortListInsert(link, 2)
    PrintList(link)
}

func SortListInsert(l *NodeI, data_ref int) *NodeI{
    if l == nil {
        return nil
    }
    n := &NodeI{Data: data_ref}
    iterator := l
    for iterator.Next != nil {
        if iterator.Data > data_ref {
            n.Next = iterator
            return n
        }
        iterator = iterator.Next
    }
    iterator.Next = n
    return l
}