package main //declarar un package al ser el principal es el main

import (
	"fmt" //paquete fmt
)

func main() { //funcion principal
	//area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println(areaCuadrado)

	x :=10
	y := 50

	//suma
	result := x + y
	fmt.Println("Suma:", result)

	//resta
	result = y - x
	fmt.Println("Resta:", result)

	//multiplicacion
	result = x * y
	fmt.Println("Multiplicación:", result)

	//division
	result = y / x
	fmt.Println("División:", result)

	//modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//incrementar
	x++
	fmt.Println("Incremental:", x)

	//decrementar
	x--
	fmt.Println("Decremental:", x)
}

