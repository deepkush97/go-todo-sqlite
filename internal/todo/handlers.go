package todo

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	todoRoutes := router.Group("/api/v1/todos")
	{
		todoRoutes.GET("/", GetTodosHandler)
		todoRoutes.GET("/:id", GetTodoByIDHandler)
		todoRoutes.POST("/", CreateTodoHandler)
		todoRoutes.PUT("/:id/complete", CompleteTodoHandler)
		todoRoutes.PUT("/:id", UpdateTodoHandler)
		todoRoutes.DELETE("/:id", DeleteTodoHandler)
	}
}

func GetTodosHandler(c *gin.Context) {
	log.Print("Fetching all todos")

	todos, err := GetTodos()
	if err != nil {
		log.Printf("Failed to fetch todos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch todos"})
		return
	}

	log.Printf("Fetched %d todos", len(todos))
	c.JSON(http.StatusOK, todos)
}

func GetTodoByIDHandler(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := GetTodoByID(id)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodoHandler(c *gin.Context) {
	log.Print("Creating a new todo")

	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		log.Printf("Invalid request payload: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := CreateTodo(todo.Title); err != nil {
		log.Printf("Failed to create todo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create todo"})
		return
	}

	log.Printf("Todo created successfully: %s", todo.Title)
	c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully"})
}

func UpdateTodoHandler(c *gin.Context) {
	log.Print("update an existing todo")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		log.Printf("Invalid request payload: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := UpdateTodo(id, todo.Title); err != nil {
		log.Printf("Failed to update todo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update todo"})
		return
	}

	log.Printf("Todo updated successfully: %s", todo.Title)
	c.JSON(http.StatusCreated, gin.H{"message": "Todo updated successfully"})
}

func CompleteTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	log.Printf("Marking todo as completed: ID %d", id)

	if err := CompleteTodoByID(id); err != nil {
		log.Printf("Failed to complete todo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to complete todo"})
		return
	}

	log.Printf("Todo marked as completed: ID %d", id)
	c.JSON(http.StatusOK, gin.H{"message": "Todo marked as completed"})
}

func DeleteTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	log.Printf("Deleting todo: ID %d", id)

	if err := DeleteTodoByID(id); err != nil {
		log.Printf("Failed to delete todo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete todo"})
		return
	}

	log.Printf("Todo deleted successfully: ID %d", id)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
