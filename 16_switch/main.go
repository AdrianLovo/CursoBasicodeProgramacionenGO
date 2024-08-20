package main //declarar un package al ser el principal es el main

import (
	"fmt" 		//paquete fmt
)


func main() { //funcion principal
	
	// iterar sobre una misma variable
	switch modulo := 5 % 2;modulo {
		case 0:
			fmt.Println("Es par")
		default:
			fmt.Println("Es impar")
	}	
	

	// sin codicion, cuando se va a evaluar multiples condiciones
	value := 50
	switch{
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}

}

