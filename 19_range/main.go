package main //declarar un package al ser el principal es el main

import "fmt"
import "strings"

func isPalindromo(text string){
	text = strings.ToLower(text)
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")
	}else{
		fmt.Println("No es palindromo")
	}
}

func main() { //funcion principal

	slice := []string{"hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}
	for i := range slice {
		fmt.Println(i)
	}
	for _, valor := range slice {
		fmt.Println(valor)
	}

	isPalindromo("Ama")
}

