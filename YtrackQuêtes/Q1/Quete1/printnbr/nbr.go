package main

import "github.com/01-edu/z01"

func main() {
	PrintNbr(-100)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	if n < 0{
		z01.PrintRune('-')
		n = -n
	}
	if n > 9{
		PrintNbr(n/10)
	}
	z01.PrintRune(rune(n%10 + '0'))
		
}