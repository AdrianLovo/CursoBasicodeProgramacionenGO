package main //declarar un package al ser el principal es el main

import (
	"fmt" 		//paquete fmt
	"strconv"
	"log"
)


func main() { //funcion principal
	valor1 := 2
	valor2 := 2

	if(valor1 == 1){
		fmt.Println("Es 1")
	}else{
		fmt.Println("No es 1")
	}

	// and
	if (valor1 == 1 && valor2 ==2){
		fmt.Println("Es verdad")
	}else{
		fmt.Println("No es verdad")
	}

	// or
	if(valor1== 0 || valor2 ==2){
		fmt.Println("Es verdad, OR")
	}else{
		fmt.Println("No es verdad, OR")
	}

	// convertir texto a numero
	value, err := strconv.Atoi("53")
	if(err != nil){		// indica si no tiene error
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

}

