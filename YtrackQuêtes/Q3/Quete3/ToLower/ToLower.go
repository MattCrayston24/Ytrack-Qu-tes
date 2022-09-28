package main

import "fmt"

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}

func ToLower(str string) string{
	var mot string
	for _, letter := range str{
		if letter >= 'A' && letter <= 'Z'{
			mot += string(letter+32)
		}else{
		mot += string(letter)
		}
	}
	return mot	
}