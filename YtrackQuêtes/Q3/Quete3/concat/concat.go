package main

import "fmt"

func main() {
	fmt.Println(Concat("Hello!", " How are you?"))

}

func Concat(str1 string, str2 string) string{
	if len (str1) == 0 {
		return str2
	}else if len (str2) == 0{
		return str1
	}
	return str1 + str2
}