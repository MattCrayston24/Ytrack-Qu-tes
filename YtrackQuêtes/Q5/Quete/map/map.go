package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}

func Map(f func(int) bool, a []int) []bool {
	var res []bool
	for _,n:=range a{
		res = append(res,f(n))
	}
	return res
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