package main
import "fmt"

func main(){
    fmt.Println(TrimAtoi("12345"))
    fmt.Println(TrimAtoi("str123ing45"))
    fmt.Println(TrimAtoi("012 345"))
    fmt.Println(TrimAtoi("Hello World!"))
    fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
    fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
    fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
    fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}

func TrimAtoi(s string) int {
    var sign int
    var result int
    var i int
    for i < len(s) {
        if s[i] == '-' {
            sign = -1
        } else if s[i] == '+' {
            sign = 1
        } else if s[i] >= '0' && s[i] <= '9' {
            result = result*10 + int(s[i]-'0')
        }
        i++
    }
    return result * sign
}