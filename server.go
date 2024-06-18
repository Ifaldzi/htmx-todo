package main

import (
	"html/template"
	"htmx-todo/controllers"
	"htmx-todo/models"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	todoItemtmpl := template.Must(template.ParseFiles("views/todo-item.html"))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == "GET") {
			todos := controllers.GetTodos()

			tmpl.Execute(w, struct {Todos []models.Todo} {Todos: todos})

			return;
		}
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == http.MethodPost) {
			newTodo := models.Todo{
				Id: uuid.NewString(),
				Title: r.FormValue("title"),
				Done: false,
			}

			controllers.AddTodo(newTodo)

			todoItemtmpl.Execute(w, newTodo)
		}
	})
	
	http.HandleFunc("/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == http.MethodDelete) {
			controllers.DeleteTodo(r.PathValue("id"))
		}

		if (r.Method == http.MethodPut) {
			updatedTodo := models.Todo{
				Id: r.PathValue("id"),
				Title: r.FormValue("title"),
				Done: r.FormValue("done") == "on",
			}

			controllers.UpdateTodo(updatedTodo)

			todoItemtmpl.Execute(w, updatedTodo)
		}
	})

	http.ListenAndServe(":80", nil)
}