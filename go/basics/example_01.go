package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. Creamos un mapa con datos iniciales
	inventario := map[string]int{
		"Laptops": 5,
		"Mouses":  20,
		"Teclados": 15,
	}

	// 2. Creamos un canal para comunicar el nombre del producto
	canalProductos := make(chan string)
	
	// Usamos WaitGroup para esperar a que la goroutine termine
	var wg sync.WaitGroup
	wg.Add(1)

	// 3. Consumidor: Una goroutine que recibe del canal y consulta el mapa
	go func() {
		defer wg.Done()
		for producto := range canalProductos {
			cantidad := inventario[producto]
			fmt.Printf("Procesando: %s | Stock disponible: %d\n", producto, cantidad)
		}
	}()

	// 4. Productor: Enviamos las llaves del mapa al canal
	for nombre := range inventario {
		canalProductos <- nombre
	}

	// Cerramos el canal y esperamos
	close(canalProductos)
	wg.Wait()
	
	fmt.Println("Procesamiento completado.")
}