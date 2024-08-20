package main //declarar un package al ser el principal es el main

import "fmt"

func main() { //funcion principal

	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//encontrar un valor
	value, ok := m["Jose"]	// ok nos indica si existe la llave
	value2 := m["Josep"] 	//si no existe la llave retorna 0
	fmt.Println(value, ok)	
	fmt.Println(value2)
}
