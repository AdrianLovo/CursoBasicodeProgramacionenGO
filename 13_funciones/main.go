package main //declarar un package al ser el principal es el main

import (
	"fmt" //paquete fmt
)


func main() { //funcion principal
	
	//for condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n");

	//for while
	counter := 0
	for counter < 10{
		fmt.Println(counter)
		counter++
	}

	//for forever
	counterForever:= 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}

