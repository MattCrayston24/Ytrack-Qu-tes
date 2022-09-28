package main

import "fmt"

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))

}

func IsAlpha(str string) bool{
	a := 0
	for _, letter := range str{
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') || (letter >= '0' && letter <= '9'){
			a += 1
		}
	}
	if a == len(str){
		return true
	}
	return false
}