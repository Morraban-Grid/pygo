package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// 1. MODELO DE DATOS
type Estudiante struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
	Curso  string `json:"curso"`
}

// 2. BASE DE DATOS EN MEMORIA
// Usamos sync.Mutex para evitar problemas de concurrencia al leer/escribir el mapa
var (
	estudiantes = make(map[int]Estudiante)
	proximoID   = 1
	mutex       sync.Mutex
)

func main() {
	// 4. ENRUTADOR (Usando las mejoras de enrutamiento de Go 1.22+)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /estudiantes", obtenerEstudiantes)
	mux.HandleFunc("GET /estudiantes/{id}", obtenerEstudiantePorID)
	mux.HandleFunc("POST /estudiantes", crearEstudiante)
	mux.HandleFunc("PUT /estudiantes/{id}", actualizarEstudiante)
	mux.HandleFunc("DELETE /estudiantes/{id}", eliminarEstudiante)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
	}
}

// ==========================================
// 3. CONTROLADORES (HANDLERS)
// ==========================================

// Obtener todos los estudiantes
func obtenerEstudiantes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mutex.Lock()
	listaEstudiantes := make([]Estudiante, 0, len(estudiantes))
	for _, estudiante := range estudiantes {
		listaEstudiantes = append(listaEstudiantes, estudiante)
	}
	mutex.Unlock()

	json.NewEncoder(w).Encode(listaEstudiantes)
}

// Obtener un estudiante por ID
func obtenerEstudiantePorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extraer el ID de la URL
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	estudiante, existe := estudiantes[id]
	mutex.Unlock()

	if !existe {
		http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(estudiante)
}

// Crear un nuevo estudiante
func crearEstudiante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nuevoEstudiante Estudiante
	// Decodificar el cuerpo JSON de la petición
	err := json.NewDecoder(r.Body).Decode(&nuevoEstudiante)
	if err != nil {
		http.Error(w, "Datos de entrada inválidos", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	nuevoEstudiante.ID = proximoID
	estudiantes[proximoID] = nuevoEstudiante
	proximoID++
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoEstudiante)
}

// Actualizar un estudiante existente
func actualizarEstudiante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var estudianteActualizado Estudiante
	err = json.NewDecoder(r.Body).Decode(&estudianteActualizado)
	if err != nil {
		http.Error(w, "Datos de entrada inválidos", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	_, existe := estudiantes[id]
	if !existe {
		mutex.Unlock()
		http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
		return
	}

	estudianteActualizado.ID = id
	estudiantes[id] = estudianteActualizado
	mutex.Unlock()

	json.NewEncoder(w).Encode(estudianteActualizado)
}

// Eliminar un estudiante
func eliminarEstudiante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	_, existe := estudiantes[id]
	if !existe {
		mutex.Unlock()
		http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
		return
	}

	delete(estudiantes, id)
	mutex.Unlock()

	w.WriteHeader(http.StatusNoContent) // 204 Éxito, pero sin contenido que retornar
}
