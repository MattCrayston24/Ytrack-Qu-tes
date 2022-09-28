package main

import "fmt"

func main() {
	fmt.Println(FindNextPrime(7))
	fmt.Println(FindNextPrime(4))
}

func FindNextPrime (nb int) int{
	if nb <= 0{
		return 0
	}
	for x := 2 ; x < nb ; x++{
		if nb % x == 0{
			nb = nb+1
		}
	}
	return nb
}