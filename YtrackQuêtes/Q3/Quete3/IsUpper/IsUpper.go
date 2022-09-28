package main

import "fmt"

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))

}

func IsUpper(str string) bool{
	a := 0
	if len(str) == 0{
		return false
	}
	for _, letter := range str{
		if letter >= 'A' && letter <= 'Z'{
			a += 1
		}
	}
	if a == len(str){
		return true
	}
	return false
}