package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. Canal para transmitir el área calculada
	canalArea := make(chan float64)

	// Lados de nuestro triángulo
	a, b, c := 3.0, 4.0, 5.0

	// 2. Goroutine para calcular el área (Herón)
	go func(ladoA, ladoB, ladoC float64) {
		// Cálculo del semiperímetro
		s := (ladoA + ladoB + ladoC) / 2
		
		// Teorema de Herón
		area := math.Sqrt(s * (s - ladoA) * (s - ladoB) * (s - ladoC))
		
		// Enviamos el resultado por el canal
		canalArea <- area
	}(a, b, c)

	// 3. Recibimos el valor del canal
	resultado := <-canalArea

	fmt.Printf("El área del triángulo calculada con Herón es: %.2f\n", resultado)
}