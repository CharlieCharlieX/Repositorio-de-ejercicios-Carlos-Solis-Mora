package main

import (
	"fmt"
	"strings"
)

/*
Trabajo realizado por Carlos Eduardo Solís Mora.
Carné: 2021051459
*/

func main() {
	str := "Hola mundo, mi nombre es Carlos\nEste es un mensaje\nPara comprobar funcionamiento."

	var letters, lines, words = countLetters(str)

	fmt.Println("Cantidad de líneas: ", lines)
	fmt.Println("Cantidad de palabras: ", words)
	fmt.Println("Cantidad de letras: ", letters)
}

func countLetters(example string) (int, int, int) {
	letters := 0
	lines := 1              //Empieza en 1 para contar la primera linea antes del primer caracter de salto de página.
	nonFunctionalWords := 1 //Empieza en 1 para contar la primer palabra antes del primer espacio.

	words := len(strings.Fields(example)) //Esta fue la mejor forma que encontré de contar las palabras. Abajo hay una
	// condicional que cuenta las palabras cada vez que hay un espacio, pero es
	//defectuoso, pues si hay varios espacios entre palabras da error.

	//linesAlternative := len(strings.Split(example, "\n")) //Esta es una alternativa también para contar la cantidad de lineas.

	for _, c := range example {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			letters++
		} else if c == '\n' {
			lines++
		} else if c == ' ' { //Condicional defectuosa, solo para comprobar si contando por espacios el resultado es correcto.
			nonFunctionalWords++
		}
	}
	//words=nonFunctionalWords //Descomentar si se quiere ver el resultado del conteo usando el filtro.
	//lines = linesAlternative //Descomentar en conjunto con la variable linesAlternative para ver el resultado.
	return letters, lines, words
}
