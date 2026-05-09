package main

import (
	"fmt"
)

// Definimos una estructura sencilla para nuestros datos
type Log struct {
	Nivel   string
	Mensaje string
}

func main() {
	// 1. Slice: Nuestra fuente de datos "cruda"
	eventos := []Log{
		{"INFO", "Sistema iniciado"},
		{"ERROR", "Fallo de conexión a BD"},
		{"INFO", "Usuario login: admin"},
		{"WARN", "Uso de CPU elevado"},
		{"ERROR", "Permiso denegado en /var/log"},
	}

	// 2. Canal: Para enviar logs uno por uno
	canalLogs := make(chan Log)

	// 3. Mapa: Para agrupar los logs (Severidad -> Lista de mensajes)
	// Usamos un slice como valor del mapa: map[string][]string
	reporte := make(map[string][]string)

	// Goroutine Procesadora
	done := make(chan bool)
	go func() {
		for log := range canalLogs {
			// Agregamos el mensaje al slice correspondiente dentro del mapa
			reporte[log.Nivel] = append(reporte[log.Nivel], log.Mensaje)
		}
		done <- true
	}()

	// 4. Productor: Recorremos el slice y enviamos al canal
	for _, e := range eventos {
		canalLogs <- e
	}
	close(canalLogs) // Importante cerrar para que el range termine

	<-done // Esperamos a que el proceso termine

	// 5. Resultado final: Mostramos el mapa organizado
	fmt.Println("--- Reporte de Sistema ---")
	for nivel, mensajes := range reporte {
		fmt.Printf("[%s]: %d eventos encontrados\n", nivel, len(mensajes))
		for _, msg := range mensajes {
			fmt.Printf("  - %s\n", msg)
		}
	}
}