//go:build ignore

package main

import "fmt"

func main() {
	// 1. Trabajando con SLICES
	// make([]T, longitud, capacidad)
	// Creamos un slice de enteros con longitud 3 y capacidad 5
	numeros := make([]int, 3, 5)
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println("--- Slices con make ---")
	fmt.Printf("Slice: %v, Longitud: %d, Capacidad: %d\n", numeros, len(numeros), cap(numeros))

	// 2. Trabajando con MAPS
	// make(map[Clave]Valor)
	// Si no usas make, el map es 'nil' y dará error al intentar asignarle datos
	estudiantes := make(map[string]int)
	estudiantes["Alice"] = 25
	estudiantes["Bob"] = 30

	fmt.Println("\n--- Maps con make ---")
	fmt.Println("Mapa de estudiantes:", estudiantes)
	fmt.Println("Edad de Alice:", estudiantes["Alice"])

	// 3. Trabajando con CHANNELS (Concurrencia)
	// make(chan T, buffer)
	// Creamos un canal para pasar strings con un buffer de 2
	mensajes := make(chan string, 2)
	mensajes <- "Hola"
	mensajes <- "Golang"

	fmt.Println("\n--- Channels con make ---")
	fmt.Println("Mensaje 1:", <-mensajes)
	fmt.Println("Mensaje 2:", <-mensajes)

	//Creando otros slices
	candidatos := make([]string, 0, 36)
	candidatos = append(candidatos,
	"karol", 
	"carlos",
	"ricardo",
	"roberto",
	"rafael",
	"cesar")
	
	//porcentaje de intencion de voto
	porcentaje := make([]float64, 0, 36)
	porcentaje = append(porcentaje, 20.5, 15.3, 10.0, 25.0, 18.2, 11.0)

	fmt.Println("\n--- Candidatos y Porcentajes ---")
	for i := range candidatos {
		fmt.Printf("%s: %.1f%%\n", candidatos[i], porcentaje[i])
	}
}
