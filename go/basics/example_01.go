package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Producto representa el modelo de datos
type Producto struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}

// Simulamos nuestra base de datos en memoria
var bd = make(map[int]Producto)
var proximoID = 1

func main() {
	// Precargamos un par de productos para probar
	bd[1] = Producto{ID: 1, Nombre: "Manzana", Precio: 1.5}
	bd[2] = Producto{ID: 2, Nombre: "Plátano", Precio: 0.8}
	proximoID = 3

	// Definimos la ruta principal del CRUD
	http.HandleFunc("/productos", manejadorProductos)
	http.HandleFunc("/productos/", manejadorProductoIndividual)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Maneja operaciones en la lista completa: GET (Leer todos) y POST (Crear)
func manejadorProductos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET": // --- LEER (TODOS) ---
		lista := []Producto{}
		for _, p := range bd {
			lista = append(lista, p)
		}
		json.NewEncoder(w).Encode(lista)

	case "POST": // --- CREAR ---
		var nuevo Producto
		err := json.NewDecoder(r.Body).Decode(&nuevo)
		if err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}
		nuevo.ID = proximoID
		bd[proximoID] = nuevo
		proximoID++

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(nuevo)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// Maneja operaciones con un ID específico: GET (Uno solo), PUT (Actualizar), DELETE (Eliminar)
func manejadorProductoIndividual(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extraer el ID de la URL (ejemplo: /productos/1)
	partes := strings.Split(r.URL.Path, "/")
	if len(partes) < 3 {
		http.Error(w, "ID requerido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(partes[2])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Verificar si el producto existe
	producto, existe := bd[id]
	if !existe {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET": // --- LEER (UNO) ---
		json.NewEncoder(w).Encode(producto)

	case "PUT": // --- ACTUALIZAR ---
		var actualizado Producto
		if err := json.NewDecoder(r.Body).Decode(&actualizado); err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}
		actualizado.ID = id // Mantener el mismo ID
		bd[id] = actualizado
		json.NewEncoder(w).Encode(actualizado)

	case "DELETE": // --- ELIMINAR ---
		delete(bd, id)
		w.WriteHeader(http.StatusNoContent) // 204 Significa éxito pero sin contenido que devolver

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}