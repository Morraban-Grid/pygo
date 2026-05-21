package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Usuario representa la estructura de datos de nuestro usuario
type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

// UsuarioHandler maneja la base de datos en memoria y la concurrencia
type UsuarioHandler struct {
	mu       sync.Mutex
	usuarios []Usuario
	nextID   int
}

func NewUsuarioHandler() *UsuarioHandler {
	return &UsuarioHandler{
		usuarios: []Usuario{
			{ID: 1, Nombre: "Alice", Email: "alice@example.com"},
		},
		nextID: 2,
	}
}

// Crear o Listar usuarios (Ruta: /usuarios)
func (h *UsuarioHandler) ManejarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		h.listarUsuarios(w, r)
	case http.MethodPost:
		h.crearUsuario(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// Obtener un usuario específico (Ruta: /usuarios/{id})
func (h *UsuarioHandler) ManejarUsuarioPorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Extraer el ID de la URL (Go 1.22+)
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	for _, u := range h.usuarios {
		if u.ID == id {
			json.NewEncoder(w).Encode(u)
			return
		}
	}

	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}

func (h *UsuarioHandler) listarUsuarios(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	// Devolvemos la lista completa
	json.NewEncoder(w).Encode(h.usuarios)
}

func (h *UsuarioHandler) crearUsuario(w http.ResponseWriter, r *http.Request) {
	var nuevoUsuario Usuario

	// Decodificar el cuerpo de la petición
	err := json.NewDecoder(r.Body).Decode(&nuevoUsuario)
	if err != nil || nuevoUsuario.Nombre == "" || nuevoUsuario.Email == "" {
		http.Error(w, "Datos de usuario inválidos", http.StatusBadRequest)
		return
	}

	h.mu.Lock()
	nuevoUsuario.ID = h.nextID
	h.nextID++
	h.usuarios = append(h.usuarios, nuevoUsuario)
	h.mu.Unlock()

	// Responder con el usuario creado y estado 201 Created
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoUsuario)
}

func main() {
	handler := NewUsuarioHandler()

	// Definición de rutas (Sintaxis compatible con Go 1.22+)
	http.HandleFunc("GET /usuarios", handler.ManejarUsuarios)
	http.HandleFunc("POST /usuarios", handler.ManejarUsuarios)
	http.HandleFunc("GET /usuarios/{id}", handler.ManejarUsuarioPorID)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
