// frequency_test.go

package main

// Importamos el paquete testing para escribir pruebas unitarias
import (
	"testing"
)

/*
Función de prueba para CountFrequency.
Aquí pasamos como argumento un slice de strings
con productos vendidos y verificamos que la función
CountFrequency devuelve el mapa correcto con la frecuencia de cada producto.
Usamos * testing.T para reportar errores en caso de que la función no devuelva el resultado esperado.
T es una estructura de tipo testing T que se usa para pruebas unitarias en Go.
*/
func TestCountFrequency(t *testing.T) {
	// Creamos un slice de productos vendidos
	ventas := []string{"manzana", "banana", "manzana", "naranja", "banana", "manzana"}

	// llamamos a la función CountFrequency con el slice de ventas
	frecuencia := CountFrequency(ventas)

	// creamos un slice de estructuras para almacenar las claves y los valores esperados
	test := []struct {
		key      string
		expected int
	}{
		{"manzana", 3},
		{"banana", 2},
		{"naranja", 1},
	}

	for _, tt := range test {
		// verificamos que la frecuencia de cada producto sea igual a la esperada
		if frecuencia[tt.key] != tt.expected {
			t.Errorf("CountFrequency(%v) = %v, expected %v", ventas, frecuencia[tt.key], tt.expected)
		}
	}
}
