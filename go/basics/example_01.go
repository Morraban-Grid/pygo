package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// User represents a user in our system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
}

// In-memory storage
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
	{ID: 3, Name: "Bob Wilson", Email: "bob@example.com", Age: 35},
}
var nextID = 4

func main() {
	// 1. Inicializamos el router por defecto de Gin
	router := gin.Default()

	// 2. Definimos las rutas y las mapeamos a sus respectivos handlers
	// Ojo: La ruta /users/search debe definirse ANTES o de forma independiente
	// para evitar colisiones con /users/:id en ciertos routers, aunque Gin maneja bien esto.
	router.GET("/users", getAllUsers)
	router.GET("/users/search", searchUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	// 3. Iniciamos el servidor en el puerto 8080
	router.Run(":8080")
}

// getAllUsers handles GET /users
func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    users,
	})
}

// getUserByID handles GET /users/:id
func getUserByID(c *gin.Context) {
	// Extraemos el parámetro ':id' de la URL (viene como string)
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Convertimos string a entero
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid ID format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	user, _ := findUserByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   "User not found",
			Code:    http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    user,
	})
}

// createUser handles POST /users
func createUser(c *gin.Context) {
	var newUser User

	// Intentamos deserializar el cuerpo JSON del Request en la estructura newUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request body",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Validamos los campos obligatorios
	if err := validateUser(newUser); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Asignamos el ID auto-incremental y actualizamos el contador global
	newUser.ID = nextID
	nextID++

	// Guardamos el nuevo usuario en el slice en memoria
	users = append(users, newUser)

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    newUser,
		Message: "User created successfully",
	})
}

// updateUser handles PUT /users/:id
func updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid ID format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Buscamos si el usuario existe antes de actualizar
	userToUpdate, index := findUserByID(id)
	if userToUpdate == nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   "User not found",
			Code:    http.StatusNotFound,
		})
		return
	}

	var updateData User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request body",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := validateUser(updateData); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Mantener el ID original de la URL y actualizar los datos en el slice
	updateData.ID = id
	users[index] = updateData

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    updateData,
		Message: "User updated successfully",
	})
}

// deleteUser handles DELETE /users/:id
func deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid ID format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	_, index := findUserByID(id)
	if index == -1 {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   "User not found",
			Code:    http.StatusNotFound,
		})
		return
	}

	// Eliminamos usando la técnica de rebanado y desempaquetado (...) que aprendimos antes
	users = append(users[:index], users[index+1:]...)

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "User deleted successfully",
	})
}

// searchUsers handles GET /users/search?name=value
func searchUsers(c *gin.Context) {
	// Capturamos el query param '?name=' de la URL
	queryName := c.Query("name")
	queryNameLower := strings.ToLower(queryName)

	var matchedUsers []User

	// Iteramos buscando coincidencias parciales (case-insensitive)
	for _, user := range users {
		if strings.Contains(strings.ToLower(user.Name), queryNameLower) {
			matchedUsers = append(matchedUsers, user)
		}
	}

	// Aunque no haya resultados, la respuesta es exitosa con un slice vacío
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    matchedUsers,
	})
}

// Helper function to find user by ID
func findUserByID(id int) (*User, int) {
	// Iteramos usando el índice para retornar la referencia directa en memoria
	for i := range users {
		if users[i].ID == id {
			return &users[i], i
		}
	}
	return nil, -1
}

// Helper function to validate user data
func validateUser(user User) error {
	// Trimmed de espacios vacíos para evitar strings con puros espacios
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email is required")
	}

	// Validación básica de formato de correo (que contenga al menos un '@' y un '.')
	if !strings.Contains(user.Email, "@") || !strings.Contains(user.Email, ".") {
		return errors.New("invalid email format")
	}

	return nil
}