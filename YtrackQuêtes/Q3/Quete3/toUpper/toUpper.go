package main

import "fmt"

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}

func ToUpper(str string) string{
	var mot string 
	for _, letter := range str{
		if letter >= 'a' && letter <= 'z'{
			mot += string(letter-32)
		}else{
		mot += string(letter)
		}
	}
	return mot
}