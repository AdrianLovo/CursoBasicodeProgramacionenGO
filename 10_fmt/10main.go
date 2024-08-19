package main //declarar un package al ser el principal es el main

import (
	"fmt" //paquete fmt
)

func main() { //funcion principal
	//declaracion de variables

	helloMessage := "Hello"
	worldMessage := "world"

	//println
	fmt.Println(helloMessage, worldMessage);	//Agrega salto de linea de forma automatica

	//printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos \n", nombre, cursos)	//%s cadena y %d numero
	fmt.Printf("%v tiene más de %v cursos \n", nombre, cursos)	//%v si no sabes el tipo de dato

	//sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	//tipo de dato
	fmt.Printf("helloMessage: %T \n", helloMessage)	//Saber el tipo de variable
	fmt.Printf("cursos: %T \n", cursos)
}

