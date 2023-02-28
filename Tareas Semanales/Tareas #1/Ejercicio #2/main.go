package main

import "fmt"

/*
Trabajo realizado por Carlos Eduardo Solís Mora.
Carné: 2021051459
*/

func main() {

	// Tamaño del centro del rombo
	centro := 13

	if centro%2 == 0 || centro <= 0 {
		fmt.Println("El tamaño del centro dado no cumple con los criterios establecidos. (El número debe ser un impar positivo).")
	} else {
		//En este primer ciclo for se crea la parte superior del rombo, se genera de forma ascendente hasta que la cantidad
		//de asteriscos impresos es la misma que la dada.
		for i := 1; i <= centro; i += 2 {
			for j := 1; j <= centro-i; j += 2 {
				fmt.Print(" ")
			}
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println(" ")
		}
		//En este segundo ciclo for se genera la parte inferior del rombo, se genera de forma descendente hasta que la cantidad
		//de asteriscos impresos sea uno. Se empieza por una cantidad de asteriscos equivalente a lo que es:
		//(Cantidad de asteriscos dados)-1 ya que en la impresión de la parte superior ya se imprimió el centro, por lo que
		//se debe comenzar a imprimir desde la siguiente iteración.
		for i := centro - 2; i >= 1; i -= 2 {
			for j := 1; j <= centro-i; j += 2 {
				fmt.Print(" ")
			}
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println(" ")
		}
	}

}
