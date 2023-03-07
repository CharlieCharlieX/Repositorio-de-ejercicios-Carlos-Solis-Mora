package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"sort"
)

/*
Trabajo realizado por Carlos Eduardo Solís Mora.
Carné: 2021051459
*/

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var resultado = l.buscarProducto(nombre)
	if resultado == -1 {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	} else {
		(*l)[resultado].cantidad = (*l)[resultado].cantidad + cantidad
		if (*l)[resultado].precio != precio {
			(*l)[resultado].precio = precio
		}
	}
}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cantidad_vender int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 {
		(*l)[prod].cantidad = (*l)[prod].cantidad - cantidad_vender
		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar
		if (*l)[prod].cantidad <= 0 {
			*l = append((*l)[:prod], (*l)[prod+1:]...)
		}
	}
}
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	var listaMinimos listaProductos
	for i := 0; i < len(*l); i++ {
		if (*l)[i].cantidad <= existenciaMinima {
			listaMinimos = append(listaMinimos, (*l)[i])
		}
	}
	return listaMinimos
}

/*
Esta es la función que se encarga de sumarle a los productos del inventario que no cumplan con la mínima cantidad requerida,
toma como parámetro la lista de esos productos que no la cumplen, y con ella recorre el slice que corresponde al inventario
para poder realizar la agregación de la cantidad faltante. Utiliza la función agregarProducto
*/
func (l *listaProductos) aumentarInventarioMinimos(listaMinimos listaProductos) listaProductos {
	for i := 0; i < len(*l); i++ {
		var producto1 = (*l)[i]
		if slices.Contains(listaMinimos, producto1) {
			l.agregarProducto(producto1.nombre, existenciaMinima-producto1.cantidad, producto1.precio)
		}
	}
	return *l
}

/*
Método para ordenar la lista lProductos según el criterio elegido, si se escribe en el string "nombre", se ordenará bajo
ese criterio, si es "cantidad" o "precio" de igual forma. La función utiliza las funcionaldiades de la librería 'sort',
las cuales modifican la lista sobre la cual se está trabajando, y no crea una nueva solo para la resolución del problema,
por lo cual, los requerimientos del enunciado se cumplen en su totalidad.
*/
func (l *listaProductos) ordenarLista(criterio string) listaProductos {
	if criterio == "nombre" {
		sort.Slice(lProductos, func(i, j int) bool {
			return lProductos[i].nombre < lProductos[j].nombre
		})
	} else if criterio == "cantidad" {
		sort.Slice(lProductos, func(i, j int) bool {
			return lProductos[i].cantidad < lProductos[j].cantidad
		})
	} else if criterio == "precio" {
		sort.Slice(lProductos, func(i, j int) bool {
			return lProductos[i].precio < lProductos[j].precio
		})
	} else {
		fmt.Println("La lista no contiene estructuras con ese atributo. No se puede ordenar")
	}
	return *l
}

func main() {
	llenarDatos()
	fmt.Println("--------------Lista Original-------------------")
	fmt.Println(lProductos)

	fmt.Println()

	fmt.Println("-----Lista luego de aumentar las cantidades hasta el mínimo establecido------")

	//Esta es la llamada a la función que aumenta
	fmt.Println(lProductos.aumentarInventarioMinimos(lProductos.listarProductosMínimos()))

	fmt.Println()

	fmt.Println("--------------Lista ordenada según criterio elegido-----------------------")

	//Aquí se aplica el método para ordenar la lista lProductos según el criterio elegido.
	fmt.Println(lProductos.ordenarLista("precio")) //Si se quiere ordenar por las otros criterios, escribir o nombre, cantidad o precio.
}
