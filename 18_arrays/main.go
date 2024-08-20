package main //declarar un package al ser el principal es el main

import "fmt"


func main() { //funcion principal
	//array (Son inmutables, no se pueden agregar mas elementos de los que estan declarados)
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//slice	no esta limitado a la cantidad de valores en instancia
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//metodos en slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])		//imprime hasta la 3 posicion (empezando en 0)
	fmt.Println(slice[2:4])		//imprime desde el 2(incluido) hasta el 3(no lo incluye)
	fmt.Println(slice[4:])		//desde el 4 en adelante (incluye el 4)

	//append
	slice = append(slice, 7)
	fmt.Println(slice, len(slice))

	//append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)	//copia los elementos de newSlice a slice
	fmt.Println(slice, len(slice))
}

