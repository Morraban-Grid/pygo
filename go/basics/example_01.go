package main

import (
	"fmt"
	"time"
)

// Definimos la estructura del Email
type Email struct {
	Destinatario string
	Asunto       string
	Cuerpo       string
}

func main() {
	// 1. Creamos el canal para transportar objetos de tipo Email
	bandejaSalida := make(chan Email)

	// 2. Iniciamos el "Servidor de Envío" como una Goroutine
	go servidorDeEnvio(bandejaSalida)

	// 3. Simulamos el envío de varios correos desde el hilo principal
	correos := []Email{
		{"profesor@universidad.edu", "Tarea Backend", "Hola, adjunto mi proyecto en Go."},
		{"amigo@gmail.com", "Fútbol", "¡Mañana jugamos a las 7!"},
		{"jefe@empresa.com", "Reporte Mensual", "El reporte de Mayo está listo."},
	}

	for _, e := range correos {
		fmt.Printf(">> Enviando correo a la bandeja de salida: %s...\n", e.Destinatario)
		bandejaSalida <- e // Enviamos el email al canal
		time.Sleep(500 * time.Millisecond) // Pequeña pausa entre clics de "Enviar"
	}

	// Cerramos el canal para avisar que no habrá más correos
	close(bandejaSalida)

	// Esperamos un poco para ver los últimos logs del servidor antes de cerrar el programa
	time.Sleep(2 * time.Second)
	fmt.Println("Proceso de usuario finalizado.")
}

// Esta función corre en segundo plano (Goroutine)
func servidorDeEnvio(entrada <-chan Email) {
	for email := range entrada {
		// Simulamos la latencia de red al enviar un email real
		fmt.Printf("   [SERVIDOR]: Procesando envío para %s...\n", email.Destinatario)
		time.Sleep(1 * time.Second) 
		fmt.Printf("   [SERVIDOR]: ¡Correo enviado con éxito! Asunto: %s\n", email.Asunto)
	}
	fmt.Println("   [SERVIDOR]: Todos los correos han sido procesados. Apagando...")
}