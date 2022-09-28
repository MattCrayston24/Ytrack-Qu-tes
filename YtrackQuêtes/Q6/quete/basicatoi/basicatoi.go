package main

import "fmt"

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string) int {
	var result int
	for i := 0 ; i < len(s) ; i++{
		if s[i] >= '0' && s[i] <= '9' {
            result = result * 10 + int(s[i]-'0')
		}
	}
	return result 
}
