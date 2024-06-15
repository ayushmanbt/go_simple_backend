package main

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/teris-io/shortid"
)

type Todo struct {
	ID     string `json:"id"`
	Desc   string `json:"desc"`
	IsDone bool   `json:"isDone"`
}

// function to create a todo from descriptio
func CreateNewTodo(desc string) (Todo, error) {

	t := new(Todo)
	//try creating id for todo

	sid, err := shortid.Generate()

	if err != nil {
		return *t, err
	}

	t.Desc = desc
	t.ID = sid
	t.IsDone = false

	return *t, nil
}

func (t *Todo) updateTodo(desc string, isDone bool) {
	t.Desc = desc
	t.IsDone = isDone
}

// intermediate type used to parse user input
type CreateTodoIntermediate struct {
	Desc string `json:"desc"`
}

// intermediate type to extract the id from delete request
type DeleteTodoIntermediate struct {
	ID string `json:"id"`
}

var todos = []Todo{}

func getTodos(context *gin.Context) {
	context.JSON(http.StatusOK, todos)
}

func getTodoById(context *gin.Context) {
	id, _ := context.Params.Get("id")
	for _, todo := range todos {
		if todo.ID == id {
			context.JSON(http.StatusOK, todo)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"messgae": fmt.Sprintf("This ID - %s does not exist", id)})
}

func createTodo(context *gin.Context) {

	//get the todo from the context
	var todoInput CreateTodoIntermediate

	if err := context.BindJSON(&todoInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad format of todo sent"})
		return
	}

	//create the final todo struct
	todo, err := CreateNewTodo(todoInput.Desc)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error generating new todo in the database", "error": err})
		return
	}

	todos = append(todos, todo)
	context.JSON(http.StatusCreated, todo)
}

func updateTodo(context *gin.Context) {
	// get the updated todo from the context
	var todoToUpdate Todo

	if err := context.BindJSON(&todoToUpdate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad format of todo sent"})
		return
	}

	for i, todo := range todos {
		if todo.ID == todoToUpdate.ID {
			todo.updateTodo(todoToUpdate.Desc, todoToUpdate.IsDone)
			todos[i] = todo
			context.JSON(http.StatusOK, todo)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("TODO with id %s was not found", todoToUpdate.ID)})
}

func deleteTodo(context *gin.Context) {
	//get the todo to be deleted
	var todoToDelete DeleteTodoIntermediate

	if err := context.BindJSON(&todoToDelete); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad format of todo sent"})
		return
	}

	newTodos := []Todo{}

	for _, todo := range todos {
		if todo.ID != todoToDelete.ID {
			newTodos = append(newTodos, todo)
		}
	}

	todos = newTodos

	context.JSON(http.StatusOK, newTodos)

}

func main() {
	router := gin.Default()
	router.GET("/todos/all", getTodos)
	router.GET("/todos/:id", getTodoById)

	router.POST("/todos/create", createTodo)

	router.PUT("/todos/update", updateTodo)

	router.DELETE("/todos/delete", deleteTodo)

	router.Run("localhost:2211")
}
