package controllers

import (
	"encoding/json"
	"fmt"
	"htmx-todo/models"
	"os"
)

func GetTodos() []models.Todo {
	content, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}

	todos := []models.Todo{}
	err = json.Unmarshal(content, &todos)
	if err != nil {
		fmt.Println(err)
	}

	return todos
}

func AddTodo(todo models.Todo) {
	todos := GetTodos()
	newTodos := append(todos, models.Todo{})
	copy(newTodos[1:], newTodos)
	newTodos[0] = todo

	saveTodos(newTodos)
}

func UpdateTodo(updatedTodo models.Todo) {
	todos := GetTodos()
	for index, todo := range todos {
		if (todo.Id == updatedTodo.Id) {
			todos[index].Title = updatedTodo.Title
			todos[index].Done = updatedTodo.Done
		}
	}

	saveTodos(todos)
}

func DeleteTodo(id string) {
	todos := GetTodos()
	indexToDelete := -1
	for index, todo := range todos {
		if (todo.Id == id) {
			indexToDelete = index
		}
	}

	if (indexToDelete != -1) {
		todos = append(todos[:indexToDelete], todos[indexToDelete+1:]...)
	
		saveTodos(todos)
	}
}

func saveTodos(newTodos []models.Todo) {
	content, err := json.Marshal(newTodos)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("data.json", content, 0644)

	if (err != nil) {
		fmt.Println(err)
	}
}