package main

import "fmt"

func main() {
	fmt.Println(IterativePower(4, 6))
}

func IterativePower(nb int, power int) int{	
	n := 1
	if power <= 0{
		return 0
	}
	for x := 1 ; x <= power ; x++{
		n = n * nb
	}
	return n
}