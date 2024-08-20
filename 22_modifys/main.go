package main //declarar un package al ser el principal es el main

import (
	mypackage "22_modifys/mypackage"
	"fmt"
)



func main() { //funcion principal

	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)
}
