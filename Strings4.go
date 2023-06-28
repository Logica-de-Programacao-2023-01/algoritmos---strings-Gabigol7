package main

import (
	"fmt"
)

func main() {
	var str1, str2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	if str1 == str2 {
		fmt.Println("As strings são iguais.")
	} else {
		fmt.Println("As strings são diferentes.")
	}
}
