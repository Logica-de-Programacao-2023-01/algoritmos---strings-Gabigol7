package main

import (
	"fmt"
	"sort"
	"strings"
)

func areAnagrams(str1, str2 string) bool {
	// Remover espaços em branco e converter para letras minúsculas
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// Verificar se as strings têm o mesmo comprimento
	if len(str1) != len(str2) {
		return false
	}

	// Converter as strings em slices de caracteres
	str1Chars := strings.Split(str1, "")
	str2Chars := strings.Split(str2, "")

	// Ordenar as slices de caracteres
	sort.Strings(str1Chars)
	sort.Strings(str2Chars)

	// Verificar se as slices ordenadas são iguais
	for i := 0; i < len(str1Chars); i++ {
		if str1Chars[i] != str2Chars[i] {
			return false
		}
	}

	return true
}

func main() {
	// Obter as strings do usuário
	var str1, str2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	// Verificar se são anagramas
	if areAnagrams(str1, str2) {
		fmt.Println("As strings são anagramas!")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}
