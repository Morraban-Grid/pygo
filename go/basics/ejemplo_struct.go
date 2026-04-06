package main

import "fmt"

func main(){
	//Definimos una estructura para representar a una persona
	type Persona struct{
		//escribimos los nombres con minuscula para que sean publicos
		nombre string
		edad int
	}

	pablo := Persona{
		nombre: "Pablo",
		edad: 30,
	}

	pablo.edad = envejarEdad(pablo.edad)

	fmt.Println("La edad de Pablo es: ", pablo.edad)



}

func envejarEdad(edad int) int{
	return edad + 10
}	