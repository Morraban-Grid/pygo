//go:build ignore
package main

import (
	"fmt"
)

func main() {
	// Definimos dos matrices de ejemplo usando slices
	matrixA := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	matrixB := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	result, err := multiplyMatrices(matrixA, matrixB)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Resultado de la multiplicación:")
	for _, row := range result {
		fmt.Println(row)
	}
}

// multiplyMatrices realiza la operación A x B
func multiplyMatrices(a, b [][]int) ([][]int, error) {
	rowsA, colsA := len(a), len(a[0])
	rowsB, colsB := len(b), len(b[0])

	// Regla fundamental: columnas de A deben ser igual a filas de B
	if colsA != rowsB {
		return nil, fmt.Errorf("dimensiones incompatibles: colsA %d != rowsB %d", colsA, rowsB)
	}

	// Inicializamos la matriz resultante con ceros
	// Tendrá el número de filas de A y el número de columnas de B
	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	// Algoritmo de multiplicación (O(n^3))
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return result, nil
}