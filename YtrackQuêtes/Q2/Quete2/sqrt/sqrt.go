package main

import "fmt"

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}

func Sqrt(nb int) int {
	a := 1
	if nb <=0{
		return 0
	}
	for x := 1 ; x < nb ; x++{
		a = x*x
		if a == nb{
			return x
		}
	}
	return 0
}