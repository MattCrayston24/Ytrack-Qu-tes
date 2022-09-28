package main

import "fmt"

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

func SortWordArr(a []string) []string{
	for i :=0 ; i < len(a) - 1 ; i++{
		for index,_:=range a{
			if index == len(a)-1{
				break
			}else if a[index] > a[index +1]{
				a[index],a[index+1] = a[index+1],a[index]
			}
		}
	}
	return a
}