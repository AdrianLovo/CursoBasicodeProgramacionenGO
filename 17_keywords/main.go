package main //declarar un package al ser el principal es el main

import (
	"fmt" 		//paquete fmt

)


func main() { //funcion principal
	//defer
	defer fmt.Println("Hola")		//ejecuta al final, un cierre de BD, cirre de archivo 
	fmt.Println("Mundo")

	//continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if(i == 2){
			fmt.Println("Es 2")
			continue		//puede ser algo que intersa que continue, aunque suceda un error debe seguir el ciclo for
		}
		if(i == 5){
			fmt.Println("Break")
			break		//cuando se cumple una condicion y se termine la ejecución
		}
	}

}

