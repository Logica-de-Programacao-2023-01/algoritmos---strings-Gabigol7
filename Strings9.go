package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string
	var oldChar, newChar string

	// Solicita a string ao usuário
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	// Solicita a letra antiga ao usuário
	fmt.Print("Digite a letra a ser substituída: ")
	fmt.Scanln(&oldChar)

	// Solicita a letra nova ao usuário
	fmt.Print("Digite a letra de substituição: ")
	fmt.Scanln(&newChar)

	// Substitui todas as ocorrências da letra antiga pela nova
	newString := strings.ReplaceAll(inputString, oldChar, newChar)

	// Imprime a nova string
	fmt.Println("Nova string:", newString)
}
