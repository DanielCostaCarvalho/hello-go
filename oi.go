package main

import "fmt"

func Oi(nome string) string {
	return "Deu bom, " + nome + "?"
}

func main() {
	fmt.Println(Oi(""))
}
