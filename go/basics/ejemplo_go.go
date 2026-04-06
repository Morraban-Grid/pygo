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

	//Creamos la matriz resultado C con las mismas dimensiones que A y B
	//Esta matriz es de tipo int
	C := make([][]int, filas)

	/*
	Vamos a iterar sobre cada fila de C
	y para cada i vamos a crear un slice de tipo int
	con la misma cantidad de columnas de A y B.
	*/
	for i := range C{
		C[i] = make([]int, columnas)
	}

	//Usamos WaitGroup para esperar a que todas las goroutines terminen
	var wg sync.WaitGroup

	//Iteramos sobre cada elemento de la matriz C
	for i := 0; i < filas; i++{
		//Avisamos que hay una nueva tarea
		wg.Add(1)

		//Lanzamos una goroutine para cada calcular cada fila
		go func(i int){
			//Al terminar la función, avisamos que la tarea ha terminado
			defer wg.Done()
			for j := 0; j < columnas; j++{
				c[i][j] = A[i][j] + B[i][j]
		}
	}(i)
	//Pasamos i como argumento para evitar problemas de concurrencia con la variable i

	}

	wg.Wait() //Esperamos a que todas las goroutines terminen
	fmt.Println("Matriz A:", A)
	fmt.Println("Matriz B:", B)
	fmt.Println("Matriz C (Resultado):", C)

}