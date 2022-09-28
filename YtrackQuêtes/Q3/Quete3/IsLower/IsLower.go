package main

import "fmt"

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}

func IsLower(str string) bool{
	a := 0
	if len(str) == 0{
		return false
	}
	for _, letter := range str{
		if letter >= 'a' && letter <= 'z'{
			a += 1
		}
	}
	if a == len(str){
		return true
	}
	return false
}