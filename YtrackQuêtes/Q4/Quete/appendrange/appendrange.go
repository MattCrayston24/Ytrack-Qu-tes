package main

import "fmt"

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}

func AppendRange(min, max int) []int{
	var a []int 
	if min < max{
		for i := min ; i < max ; i++{
			a = append(a, i)
		}
	}
	return a
}