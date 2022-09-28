package main

func main() {

	a := []string{"Hello", "how", "are", "you?"}
	PrintWordsTables(a)
}

func PrintWordsTables(a []string) {
	for _, lettre := range a{
		print (lettre, "\n")
	}
}