package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var oldChar string
	var newChar string

	// Solicitar a string ao usuário
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	// Solicitar o caractere a ser substituído
	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanln(&oldChar)

	// Solicitar o caractere de substituição
	fmt.Print("Digite o caractere de substituição: ")
	fmt.Scanln(&newChar)

	// Substituir as ocorrências do caractere na string
	newString := strings.ReplaceAll(input, oldChar, newChar)

	// Imprimir o resultado
	fmt.Println("Resultado: ", newString)
}
