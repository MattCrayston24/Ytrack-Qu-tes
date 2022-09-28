package main

import "fmt"

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}

func ConcatParams(arg []string) string{
	var anwser string
	if len(arg) > 0{
		for _, a := range arg{
			anwser += a + "\n"
		}
	}
	return anwser
}