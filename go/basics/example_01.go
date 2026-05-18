package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Estructura para la respuesta del clima
type ClimaRespuesta struct {
	Ciudad      string `json:"ciudad"`
	Temperatura string `json:"temperatura"`
	Condicion   string `json:"condicion"`
}

// Simulamos datos de clima para algunas ciudades
var datosClima = map[string]ClimaRespuesta{
	"madrid":    {Ciudad: "Madrid", Temperatura: "22°C", Condicion: "Soleado"},
	"bogota":    {Ciudad: "Bogotá", Temperatura: "14°C", Condicion: "Lluvioso"},
	"buenosaires": {Ciudad: "Buenos Aires", Temperatura: "18°C", Condicion: "Nublado"},
}

func manejadorClima(w http.ResponseWriter, r *http.Request) {
	// 1. Validar que solo se permita el método GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// 2. Extraer el parámetro "ciudad" de la URL (?ciudad=...)
	// r.URL.Query() parsea la URL y .Get() busca la clave
	ciudadSolicitada := r.URL.Query().Get("ciudad")

	// Si el usuario no envió el parámetro, devolvemos un error 400
	if ciudadSolicitada == "" {
		http.Error(w, "Falta el parámetro 'ciudad' en la URL", http.StatusBadRequest)
		return
	}

	// Limpiamos el texto (minúsculas y sin espacios) para buscar en el mapa
	ciudadClave := strings.ToLower(strings.TrimSpace(ciudadSolicitada))

	// 3. Buscar la ciudad en nuestro "mapa" (base de datos)
	clima, existe := datosClima[ciudadClave]

	w.Header().Set("Content-Type", "application/json")

	if !existe {
		// Si la ciudad no está, devolvemos un 404 en formato JSON
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Ciudad no encontrada en el sistema"})
		return
	}

	// 4. Si todo está bien, devolvemos el clima de la ciudad
	json.NewEncoder(w).Encode(clima)
}

func main() {
	// Registramos la ruta /clima
	http.HandleFunc("/clima", manejadorClima)

	port := ":8080"
	fmt.Printf("Servidor de clima corriendo en http://localhost%s\n", port)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}