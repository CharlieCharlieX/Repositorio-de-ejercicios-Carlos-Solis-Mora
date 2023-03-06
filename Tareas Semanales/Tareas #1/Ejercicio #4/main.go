package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
Trabajo realizado por Carlos Eduardo Solís Mora.
Carné: 2021051459
*/

// Estructura persona, cada persona cuenta con un nombre de tipo string, y una edad de tipo int32
type persona struct {
	nombre string
	edad   int32
}

// Lista para almacenar las personas
type Personas []persona

// Variable necesaria para la creación de personas
var humano Personas

// La constante para determinar si una persona es mayor de edad o no
const edadAdulta = 18

// Función para agregar nuevas personas a la lista Personas
func (f *Personas) agregarPersona(nom string, eda int32) {
	idx := slices.IndexFunc(*f, func(p persona) bool { return p.nombre == nom })
	if idx == -1 {
		*f = append(*f, persona{nom, eda})
	} else {
		fmt.Println("cant add existing product to the list")
	}
}

/*
	Función encargada de calcular las personas que son mayores de edad. Utiliza la función filter any, y devuelve una lista

de tipo persona para iterarla posteriormente y obtener los distintos valores de esas personas.
*/
func (f *Personas) calcularMayorEdad() []persona {
	lista := filterAny(*f, func(p persona) bool {
		return p.edad >= edadAdulta
	})
	return lista
}

// Función encargada de filtrar cualquier estructura que cumpla con los requisitos previamente establecidos.
func filterAny[P1 any](list []P1, f func(P1) bool) []P1 {
	filtered := make([]P1, 0)

	for _, element := range list {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func main() {
	//Se crean varias personas.
	humano.agregarPersona("Pablo", 15)
	humano.agregarPersona("José", 18)
	humano.agregarPersona("María", 13)
	humano.agregarPersona("Eduardo", 20)
	humano.agregarPersona("Ricardo", 11)

	//Se calculan las personas mayores de edad y se almacenan en una lista.
	personasMayores := humano.calcularMayorEdad()

	//Se recorre la lista previamente creada, y se obtiene el nombre y la edad de las personas mayores o iguales al os 18 años.
	for _, p := range personasMayores {
		fmt.Printf("Nombre: %s, Edad: %d\n", p.nombre, p.edad)
	}

}
