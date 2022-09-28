package main

import "fmt"

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}

func IsPrime(nb int) bool {
	if nb <= 1{
		return false
	}
	for x := 2 ; x < nb ; x++{
		if nb % x == 0{
			return false

		}
	}
	return 	true

}