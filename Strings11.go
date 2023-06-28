package main

import (
	"fmt"
	"strings"
)

func removeVogais(str string) string {
	vogais := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	// Substituir vogais por uma string vazia
	for _, vogal := range vogais {
		str = strings.ReplaceAll(str, vogal, "")
	}

	return str
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	resultado := removeVogais(input)
	fmt.Println("Resultado:", resultado)
}
