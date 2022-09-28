package main

import "fmt"

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This","1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}

func CountIf(f func(string) bool, tab []string) int {
	x := 0
	for _,n:=range tab{
		if f(n) == true{
			x += 1
		}
	}
	return x
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