package main //declarar un package al ser el principal es el main

import (
	"fmt" //paquete fmt
)

func normalFUnction(message string){
	fmt.Println(message)
}

func tripleArgument(a, b int, c string){	// (a int, b int, c string) pude ser de esta forma tambien
	fmt.Println(a, b, c)	
}

func returnValue(a int) int{	//retornando 1 volor
	return a * 2
}

func dobleReturn(a int) (c, d int){		//retorna dos valores
	return a, a * 2
}

func main() { //funcion principal
	normalFUnction("Hola mundo")
	tripleArgument(1, 2, "Hola")
	value := returnValue(2)
	fmt.Println("value:", value)

	value1, value2 := dobleReturn(2)	//recibiendo dos valroes
	fmt.Println("value1 y value2",value1, value2)

	only1, _ := dobleReturn(8)	//recibiendo dos valroes
	fmt.Println("only1",only1)
}

