package main
import "fmt"
func main(){
    strs := []string{"Hello!", " How", " are", " you?"}
    fmt.Println(Join(strs, ":"))
}

func Join(strs []string, sep string) string {
    mot := ""
    for index,letter := range strs{
        if index == len(strs)-1{
            mot+=letter
        }else{
            mot+= letter + sep
        }
    }
    return mot
}