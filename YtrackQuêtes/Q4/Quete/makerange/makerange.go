package main

import "fmt"

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}

func MakeRange(min int, max int) []int{
	size := max - min
	var anwser []int
	if size > 0{
		anwser := make([]int, size)
		for i := 0 ; i < size ; i++{
			anwser[i] = i + min
		}
		return anwser
	}
	return anwser
}