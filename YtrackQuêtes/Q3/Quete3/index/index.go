package main

import "fmt"

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}

func Index (s string, i string) int{
	for index,lettre := range s{
        if lettre == rune(i[0]){
            for index2,lettre2 := range i{
                if lettre2 != rune(s[index+index2]){
                    break
                }else if index2 == len(i)-1{
                    return index
                }

            }

        } 
    }
    return -1
}