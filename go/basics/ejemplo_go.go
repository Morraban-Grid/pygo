//go:build ignore
package main

import (
	"fmt"
)

type Persona struct{
		Nombre string
		Edad int
}

type Empleado struct{
		Persona
		Salario float64
}

func aumentarSalario(e *Empleado) {
		e.Salario += 5001.00
}

func main(){

	empleado := Empleado{
		Persona: Persona{
			Nombre: "Carlos",
			Edad: 30,
		},
		Salario: 50000.00,
	}
	fmt.Printf("Empleado: %s, Edad: %d, Salario: %.2f\n", empleado.Nombre, empleado.Edad, empleado.Salario)	

	aumentarSalario(&empleado)

	fmt.Printf("Nuevo salario: %.2f\n", empleado.Salario)
}



