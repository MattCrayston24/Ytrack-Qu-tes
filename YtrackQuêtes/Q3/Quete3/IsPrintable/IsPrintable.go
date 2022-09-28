package main

import "fmt"

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))

}

func IsPrintable(str string) bool{
	a := 0
	if len(str) == 0{
		return false
	}
	for _, letter := range str{
		if letter > 32 && letter < 126{
			a += 1
		}
	}
	if a == len(str){
		return true
	}
	return false
}