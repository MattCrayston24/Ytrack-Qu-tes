package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}

func SplitWhiteSpaces( a string) []string{
	var ch []string
	mot := ""
	for _, letter := range  a{
		if string(letter) != " "{
			mot += string(letter)
		}else{
			ch = append(ch, mot)
			mot = ""
		}
			
	}
	if mot != ""{
		ch = append(ch , mot)
	}
	return ch
}