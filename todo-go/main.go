package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean home", Completed: false},
	{ID: "2", Item: "Read book", Completed: false},
	{ID: "3", Item: "Watch series", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)

}

func postTodos(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {

	id := context.Param("id")
	todo, err := getById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func putTodo(context *gin.Context) {
	id := context.Param("id")
	currentTodo, err := getById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	var updatedTodo todo
	if err := context.BindJSON(&updatedTodo); err != nil {
		return
	}

	// Update only the fields that are provided in the request
	if updatedTodo.Item != "" {
		currentTodo.Item = updatedTodo.Item
	}
	if updatedTodo.Completed {
		currentTodo.Completed = updatedTodo.Completed
	}

	context.IndentedJSON(http.StatusOK, currentTodo)
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", postTodos)
	router.PATCH("todos/:id", toggleTodoStatus)
	router.PUT("/todos/:id", putTodo)
	router.Run("localhost:9090")
}
