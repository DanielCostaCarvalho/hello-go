package main

import "fmt"

func Hello(name string) string {
	if(name == "") {
		const portugueseSolo = "Deu bom?"
		return portugueseSolo
	}

	
	const portuguesePrefix = "Deu bom, "
	return portuguesePrefix + name + "?"
}

func main() {
	fmt.Println(Hello(""))
}
