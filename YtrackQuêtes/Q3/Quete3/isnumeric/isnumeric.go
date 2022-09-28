package main

import "fmt"

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
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