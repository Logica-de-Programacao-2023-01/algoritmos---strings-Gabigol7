package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if containsNumber(input) {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém números.")
	}
}

func containsNumber(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
