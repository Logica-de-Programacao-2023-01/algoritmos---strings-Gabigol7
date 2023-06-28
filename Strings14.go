package main

import (
	"fmt"
	"strconv"
)

func isDecrescente(str string) bool {
	if len(str) < 2 {
		return true
	}

	for i := 1; i < len(str); i++ {
		curr, _ := strconv.Atoi(string(str[i]))
		prev, _ := strconv.Atoi(string(str[i-1]))

		if curr >= prev {
			return false
		}
	}

	return true
}

func main() {
	var input string
	fmt.Print("Digite uma sequência numérica: ")
	fmt.Scanln(&input)

	if isDecrescente(input) {
		fmt.Println("A sequência é decrescente.")
	} else {
		fmt.Println("A sequência não é decrescente.")
	}
}
