package main

import (
	"fmt"
)

// productor: envía números y luego cierra el canal
func generarNumeros(max int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= max; i++ {
			out <- i
		}
	}()
	return out
}

// transformador: recibe números, envía su cuadrado y cierra salida
func calcularCuadrados(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func main() {
	numeros := generarNumeros(5)
	cuadrados := calcularCuadrados(numeros)

	// consumidor final
	for c := range cuadrados {
		fmt.Println(c)
	}
}