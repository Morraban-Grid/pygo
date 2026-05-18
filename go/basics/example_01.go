package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Definimos la estructura de nuestro Producto
type Producto struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Precio float64 `json:"precio"`
}

// Simulamos una base de datos en memoria con una rebanada (slice)
var productos = []Producto{
	{ID: 1, Nombre: "Laptop", Precio: 899.99},
	{ID: 2, Nombre: "Ratón Óptico", Precio: 19.99},
}

// Controlador para manejar las peticiones de /productos
func manejadorProductos(w http.ResponseWriter, r *http.Request) {
	// Definimos que la respuesta siempre será JSON
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// Convertimos los productos a JSON y los enviamos
		json.NewEncoder(w).Encode(productos)

	case http.MethodPost:
		var nuevoProducto Producto
		// Decodificamos el cuerpo de la petición (JSON) dentro de la estructura
		err := json.NewDecoder(r.Body).Decode(&nuevoProducto)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		// Asignamos un ID simple y guardamos
		nuevoProducto.ID = len(productos) + 1
		productos = append(productos, nuevoProducto)

		// Respondemos con el producto creado y el estatus 201 (Created)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(nuevoProducto)

	default:
		// Si usan PUT, DELETE, etc., respondemos que no está permitido
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Asociamos la ruta "/productos" con nuestra función manejadora
	http.HandleFunc("/productos", manejadorProductos)

	// Iniciamos el servidor en el puerto 8080
	fmt.Println("Servidor corriendo en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}