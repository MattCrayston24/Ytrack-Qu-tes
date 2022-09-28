package main

import "fmt"

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev( s string) string{
	var tab []rune
	for i := len(s) - 1; i >= 0 ; i--{
		//print (i)
		if	i >= 0{
			tab = append(tab,rune(s[i]))
			//print (tab)
		}
	}
	return string(tab)
}