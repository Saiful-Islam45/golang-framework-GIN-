package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Todo represents a todo item
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos []Todo

func main() {
	router := gin.Default()

	// Get all todos
	router.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})

	// Create a new todo
	router.POST("/todos", func(c *gin.Context) {
		var todo Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todos = append(todos, todo)
		c.JSON(http.StatusCreated, todo)
	})

	// Get a specific todo
	router.GET("/todos/:id", func(c *gin.Context) {
		id,_:=strconv.Atoi(c.Param("id"))
		for _, todo := range todos {
			if todo.ID == id {
				c.JSON(http.StatusOK, todo)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	})

	// Update a specific todo
	router.PUT("/todos/:id", func(c *gin.Context) {
		id,_:=strconv.Atoi(c.Param("id"))
		var todo Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i] = todo
				c.JSON(http.StatusOK, todo)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	})

	// Delete a specific todo
	router.DELETE("/todos/:id", func(c *gin.Context) {
		id,_:=strconv.Atoi(c.Param("id"))
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
				return
			}
		}
	})
	router.Run("localhost:8080")
}