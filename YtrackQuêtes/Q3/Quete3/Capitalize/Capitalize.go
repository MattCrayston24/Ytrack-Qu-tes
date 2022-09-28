package main

import "fmt"

func main(){
    fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s string)string{
    mot := ""
    var count rune 
    for i,lettre := range s {
        if lettre<47{
            count = lettre
        }
        if i == 0 || rune(s[i-1]) == count{
            if lettre >= 'a' && lettre <= 'z'{
                mot += string(lettre-32)
                }else {
                    mot += string(lettre)
                }
        }else {
            if lettre >= 'A' && lettre <= 'Z'{
                    mot += string(lettre+32)
            }else {
                mot += string(lettre)
            }
        }
    }
    return mot
}