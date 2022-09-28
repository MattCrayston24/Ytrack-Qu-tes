package main

import "fmt"

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func Any(f func(string) bool, a []string) bool {
	for _,n:=range a{
		if f(n) == true{
			return true
		}
	}
	return false
}

func IsNumeric(str string) bool{
	a := 0
	for _, letter := range str{
		if letter >= '0' && letter <= '9'{
			a += 1
		}
	}
	if a == len(str){
		return true
	}
	return false
}