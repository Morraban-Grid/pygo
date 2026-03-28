//go:build ignore

/*
package main

import "fmt"

func main() {

	//creando un mapa de string a int usando make
	mapa01 := make(map[string]int)
	mapa01["pedro"] = 30
	mapa01["kiara"] = 40
	mapa01["robert"] = 1

	//comprobar si el valor existe en el mapa
	edad, ok := mapa01["pablo"]
	if ok {
		fmt.Println("Edad de Pedro:", edad)
	} else {
		fmt.Println("No encontrado en el mapa")
	}

	fmt.Println("mapa: ", mapa01)
}
*/

package main

import "fmt"

func main() {

	// creamos un slice de ventas de productos
	ventas := []string{"TV", "TV", "TV", "PC", "TV", "Mouse", "PC", "TV"}
	// creamos un mapa para contar la frecuencia de cada producto
	frecuencia := make(map[string]int)

	for _, producto := range ventas {
		frecuencia[producto]++
	}

	fmt.Println(frecuencia)
	// Output: map[TV:3 PC:2 Mouse:1]
	// Nota: el orden de impresión puede variar en cada ejecución
}
