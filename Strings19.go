package main

import (
	"fmt"
)

func main() {
	var input string

	// Solicitar a string ao usuário
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	// Inverter a string
	reversed := reverseString(input)

	// Imprimir a string invertida
	fmt.Println("String invertida:", reversed)
}

func reverseString(input string) string {
	// Converter a string para um slice de runas
	runes := []rune(input)

	// Inverter o slice de runas
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Converter o slice de runas de volta para uma string
	reversed := string(runes)

	return reversed
}
