package main

import "fmt"

func main(){
    elems := []string{"Hello!", " How", " are", " you?"}
    fmt.Println(BasicJoin(elems))
}

func BasicJoin(elems []string)string{
    mot := ""
    for _,letter := range elems{
        mot+= letter
    }
    return mot
}