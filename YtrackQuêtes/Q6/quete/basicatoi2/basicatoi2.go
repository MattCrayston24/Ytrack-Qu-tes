package main

import "fmt"

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}


func BasicAtoi2(s string) int {
	var result int
	for i := 0 ; i < len(s) ; i++{
		if s[i] > 48 && s[i] <= 57{
			result = result * 10 + int(s[i]-'0')
		}else if !(s[i] >= 48 && s[i] <= 57){
			result = 0
			break
		}
	}
	return result 
}
