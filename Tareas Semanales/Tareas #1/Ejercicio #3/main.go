package main

import "fmt"

/*
Trabajo realizado por Carlos Eduardo Solís Mora.
Carné: 2021051459
*/

/*
Esta función se encarga de rotar el slice
Slice: El slice a rotar
Positions: es la cantidad de espacios en las que se debe rotar el slice
Direction: La dirección a la que se rotará el slice.
TRUE = Derecha
FALSE = Izquierda
*/
func rotar(slice []int, positions int, direction bool) {
	var pivote int //Variable para realizar el cambio de posición
	var temp int   //El valor que se cambiará de posición en el siguiente ciclo

	if direction == true { //Si direction es true, se gira a la derecha
		for j := 0; j < positions; j++ {
			for i := 0; i < len(slice)-1; i++ {
				if i == 0 {
					temp = slice[i+1]
					slice[i+1] = slice[i]
					slice[i] = slice[len(slice)-1]
				} else {
					pivote = slice[i+1]
					slice[i+1] = temp
					temp = pivote
				}
			}
		}
	} else if direction == false { //Si direction es false, se gira a la izquierda
		for j := 0; j < positions; j++ {
			for i := len(slice) - 1; i > 0; i-- {
				if i == len(slice)-1 {
					temp = slice[i-1]
					slice[i-1] = slice[i]
					slice[i] = slice[0]
				} else {
					pivote = slice[i-1]
					slice[i-1] = temp
					temp = pivote
				}
			}
		}
	}
}

func main() {
	var slice []int           // Se inicializa el slice
	slice = append(slice, 10) //Se llenan con diferentes valores
	slice = append(slice, 20)
	slice = append(slice, 30)
	slice = append(slice, 40)
	slice = append(slice, 50)
	slice = append(slice, 60)
	fmt.Println("El slice original es este: ", slice)
	rotar(slice, 1, true) //call the rotate function
	fmt.Println("El slice rotado es: ", slice)
}
