package main //declarar un package al ser el principal es el main

import (
	"fmt" //paquete fmt
)

func mian() { //funcion principal
	//declaracion de constantes
	const pi float64 = 3.14 //declarando el tipo de dato
	const pi2 = 3.1415      //sin declarar el tipo

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//declaracion de variables enteras
	base := 12 // : indican que se va a crear por primera vez y deduce el tipo que sera
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//zero values: tiene por defecto un valor 
	var a int			//0
	var b float64		//0
	var c string		//vacio
	var d bool			//false

	fmt.Println(a, b, c, d)
}

