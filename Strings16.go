package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Println("Digite a primeira string:")
	fmt.Scanln(&str1)

	fmt.Println("Digite a segunda string:")
	fmt.Scanln(&str2)

	if strings.Contains(str1, str2) {
		fmt.Println("A segunda string é uma substring da primeira.")
	} else {
		fmt.Println("A segunda string NÃO é uma substring da primeira.")
	}
}
