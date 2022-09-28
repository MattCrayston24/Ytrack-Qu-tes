package main 

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
	for index, letter := range s{
		if index == len(s) - 1{
			return letter
		}
	}
	return 0
}	
