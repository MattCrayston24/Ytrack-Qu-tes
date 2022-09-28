package main

import "fmt"

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	n := 1
	for  x := 1 ; x <= nb ; x++{
		n = n * x
	}
	return n
}