//go:build ignore

package main

import "fmt"

func main() {

	//creando un mapa de string a int usando make
	mapa01 := make(map[string]int)
	mapa01["pedro"] = 30
	mapa01["kiara"] = 40
	mapa01["robert"] = 1

	fmt.Println("Hola mundo ", mapa01)
	//fmt.Println("\n\n")

	//creando un slice de strings con longitud 3 y capacidad de 10
	slice01 := make([]string, 3, 10)
	slice01[0] = "hola"
	slice01[2] = ", esto es un programa en go."
	slice01[1] = "mundo"

	for i := 0; i < len(slice01); i++ {
		fmt.Println(slice01[i])
	}

	//creando un canal con un buffer de tamaño 4
	canal01 := make(chan string, 4)
	canal01 <- "Hola, esto es un mensaje en un canal. Usamos un goroutine para enviar este mensaje."
	canal01 <- "Este es otro mensaje en el canal."
	canal01 <- "Este es el tercer mensaje del canal."
	canal01 <- "cuarte mensaje."

	salida1 := <-canal01
	salida2 := <-canal01
	salida3 := <-canal01

	fmt.Println("\n\n--- Mensajes del canal ---")
	fmt.Println(salida1) // Imprime el primer mensaje directamente
	fmt.Println(salida2)
	fmt.Println(salida3)
	fmt.Println(<-canal01) // Imprime el segundo mensaje directamente
}
