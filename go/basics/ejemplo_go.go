//go:build ignore
package main

import (
	"fmt"
)

func main(){
	type Persona struct{
		Nombre string
		Edad int
	}

	type Empleado struct{
		Persona
		Salario float64
	}

	empleado := Empleado{ Nombre: "Carlos", Edad: 30, Salario: 50000.0
	}

	fmt.Printf("Empleado: %s, Edad: %d, Salario: %.2f\n", empleado.Nombre, empleado.Edad, empleado.Salario)	
}