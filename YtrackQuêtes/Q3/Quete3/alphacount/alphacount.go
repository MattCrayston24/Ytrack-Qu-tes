package main

import "fmt"

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

func AlphaCount(s string) int{
	nb := 0
	if len(s) <= 0{
		return 0
	}
	for _, letter := range s{
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z'){
			nb += 1
		}
	}
	return nb
}