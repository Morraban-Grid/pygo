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
	// Inicializamos el router por defecto de Gin
	router := gin.Default()

	// Setup routes
	router.GET("/users", getAllUsers)
	router.GET("/users/search", searchUsers) // Colocado antes para priorizar la coincidencia en el ruteo
	router.GET("/users/:id", getUserByID)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	// Iniciamos el servidor en el puerto 8080
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

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request body",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := validateUser(newUser); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	newUser.ID = nextID
	nextID++

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

	users = append(users[:index], users[index+1:]...)

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "User deleted successfully",
	})
}

// searchUsers handles GET /users/search?name=value
func searchUsers(c *gin.Context) {
	queryName, exists := c.GetQuery("name")

	// Si el parámetro no se envía o está completamente en blanco, lanzamos el error 400 exigido
	if !exists || strings.TrimSpace(queryName) == "" {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Missing or empty 'name' query parameter",
			Code:    http.StatusBadRequest,
		})
		return
	}

	queryNameLower := strings.ToLower(queryName)
	matchedUsers := []User{} // Inicializado explícitamente vacío para serializar [] y no null

	for _, user := range users {
		if strings.Contains(strings.ToLower(user.Name), queryNameLower) {
			matchedUsers = append(matchedUsers, user)
		}
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    matchedUsers,
	})
}

// Helper function to find user by ID
func findUserByID(id int) (*User, int) {
	for i := range users {
		if users[i].ID == id {
			return &users[i], i
		}
	}
	return nil, -1
}

// Helper function to validate user data
func validateUser(user User) error {
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email is required")
	}

	if !strings.Contains(user.Email, "@") || !strings.Contains(user.Email, ".") {
		return errors.New("invalid email format")
	}

	return nil
}
