package main

import "fmt"

const spanishPrefix = "Estudie, "
const frenchPrefix = "Bonjour, "
const portuguesePrefix = "Deu bom, "

const empty = ""

const spanishLanguage = "espanhol"
const frenchLanguage = "francês"

func Hello(name, language string) string {
	if(name == empty) {
		const portugueseSolo = "Deu bom?"
		return portugueseSolo
	}
	
	
	return getPrefix(language) + name
}

func getPrefix(language string) string {
	
	prefix := portuguesePrefix

	switch language {
		case spanishLanguage:
			prefix = spanishPrefix
		case frenchLanguage:
			prefix = frenchPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("", ""))
}
