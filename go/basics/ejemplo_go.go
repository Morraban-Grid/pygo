//go:build ignore

package main

import (
	"fmt"
	"sync"
)

//Ejemplo de goroutines
func main(){
	//matriz A, es en realidad un slice de slices, pero se puede usar como matriz
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	//matriz B
	B := [][]int{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}

	//matriz resultado C
	filas := len(A) //len(A) es el número de filas de A
	columnas := len(A[0]) //len(A[0]) es el numero de columnas de A

}