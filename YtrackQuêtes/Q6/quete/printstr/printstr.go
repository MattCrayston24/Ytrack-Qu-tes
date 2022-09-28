package main

import "fmt"

func main() {
	PrintStr("Hello World!")
}

func PrintStr(a string){
	for i, _ := range a{
		fmt.Printf(string(a[i]))
	}
}