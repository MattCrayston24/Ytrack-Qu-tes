package main

import "fmt"

func main() {
	s := "Hello4 !HAaaaaahow4 !HAaaaaaare4 !HAaaaaayou?"
	fmt.Printf("%#v\n",Split(s, "4 !HAaaaaa"))
}

func Split(s, sep string) []string {
	var tab []string
	mot := ""
	for i:= 0; i < len(s)+1 ; i++{
		if len(s[i:len(s)]) >= len(sep){
			if string(s[i:i+len(sep)]) != sep{
				mot += string(s[i]) 
				if len(s[i+1:len(s)]) < len(sep){
					mot += string(s[i+1:len(s)]) 
				}
			}else {
				tab = append(tab, mot)
				mot = ""
				i = i + len(sep)-1
			}
		}else{
			mot = string(s[i:len(s)])
			tab = append(tab, mot)
			mot = ""
			i = i + len(sep)
		
		}
	}	
	if mot != ""{
		tab = append(tab, mot)
	}
	return tab
}