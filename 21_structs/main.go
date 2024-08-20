package main //declarar un package al ser el principal es el main
import "fmt"

type car struct {
	brand string
	year  int
}

func main() { //funcion principal

	//primera forma
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//segunda forma
	var otherCar car
	otherCar.brand = "Ferrai"
	otherCar.year = 200
	fmt.Println(otherCar)
}
