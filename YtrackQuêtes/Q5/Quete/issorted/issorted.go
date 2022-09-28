package main

import "fmt"

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(Sorted, a1)
	result2 := IsSorted(Sorted, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func IsSorted(f func(a, b int) int, tab []int) bool {
	tabl := 1
	for i:=0 ; i<len(tab)-1 ; i++{	
		a := tab[i]
		b := tab[i+1]
		if f(a,b) <= 0{
			tabl += 1
		}
		if tabl == len(tab){
			return true
		}
	}
	return false
}



func Sorted(a, b int) int {
	if a > b{
		return 1
	}
	if a == b{
		return 0
	}else{
		return -1
	}
}