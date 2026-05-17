package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Tarea representa el modelo de datos
type Tarea struct {
	ID        int
	Contenido string
}

// "Base de datos" en memoria usando un slice
var tareas []Tarea
var contadorID = 1

func main() {
	// Precargamos un par de tareas (Crear)
	tareas = append(tareas, Tarea{ID: 1, Contenido: "Estudiar Go"})
	tareas = append(tareas, Tarea{ID: 2, Contenido: "Comprar café"})
	contadorID = 3

	lector := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- MENÚ CRUD DE TAREAS ---")
		fmt.Println("1. Ver tareas (Read)")
		fmt.Println("2. Agregar tarea (Create)")
		fmt.Println("3. Editar tarea (Update)")
		fmt.Println("4. Eliminar tarea (Delete)")
		fmt.Println("5. Salir")
		fmt.Print("Elige una opción: ")

		opcionRaw, _ := lector.ReadString('\n')
		opcion := strings.TrimSpace(opcionRaw)

		switch opcion {
		case "1":
			leerTareas()
		case "2":
			crearTarea(lector)
		case "3":
			actualizarTarea(lector)
		case "4":
			eliminarTarea(lector)
		case "5":
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida. Intenta de nuevo.")
		}
	}
}

// === 1. LEER (READ) ===
func leerTareas() {
	fmt.Println("\n--- LISTA DE TAREAS ---")
	if len(tareas) == 0 {
		fmt.Println("[ No hay tareas pendientes ]")
		return
	}
	for _, t := range tareas {
		fmt.Printf("[%d] %s\n", t.ID, t.Contenido)
	}
}

// === 2. CREAR (CREATE) ===
func crearTarea(lector *bufio.Reader) {
	fmt.Print("\nEscribe la nueva tarea: ")
	texto, _ := lector.ReadString('\n')
	texto = strings.TrimSpace(texto)

	if texto == "" {
		fmt.Println("La tarea no puede estar vacía.")
		return
	}

	nueva := Tarea{ID: contadorID, Contenido: texto}
	tareas = append(tareas, nueva)
	contadorID++
	fmt.Println("¡Tarea agregada con éxito!")
}

// === 3. ACTUALIZAR (UPDATE) ===
func actualizarTarea(lector *bufio.Reader) {
	leerTareas()
	if len(tareas) == 0 {
		return
	}

	fmt.Print("\nIngresa el ID de la tarea a editar: ")
	idRaw, _ := lector.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idRaw))
	if err != nil {
		fmt.Println("ID inválido.")
		return
	}

	// Buscar la tarea por ID
	indice := -1
	for i, t := range tareas {
		if t.ID == id {
			indice = i
			break
		}
	}

	if indice == -1 {
		fmt.Println("Tarea no encontrada.")
		return
	}

	fmt.Print("Escribe el nuevo contenido: ")
	nuevoTexto, _ := lector.ReadString('\n')
	nuevoTexto = strings.TrimSpace(nuevoTexto)

	tareas[indice].Contenido = nuevoTexto
	fmt.Println("¡Tarea actualizada!")
}

// === 4. ELIMINAR (DELETE) ===
func eliminarTarea(lector *bufio.Reader) {
	leerTareas()
	if len(tareas) == 0 {
		return
	}

	fmt.Print("\nIngresa el ID de la tarea a eliminar: ")
	idRaw, _ := lector.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idRaw))
	if err != nil {
		fmt.Println("ID inválido.")
		return
	}

	indice := -1
	for i, t := range tareas {
		if t.ID == id {
			indice = i
			break
		}
	}

	if indice == -1 {
		fmt.Println("Tarea no encontrada.")
		return
	}

	// Eliminar del slice manteniendo el orden
	tareas = append(tareas[:indice], tareas[indice+1:]...)
	fmt.Println("¡Tarea eliminada!")
}