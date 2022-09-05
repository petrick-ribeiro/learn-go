package main

import (
	"apis/todo-pg/configs"
	"apis/todo-pg/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main()  {
  err := configs.Load()
  if err != nil {
    panic(err)
  }

  r := chi.NewRouter()
  r.Post("/", handlers.Create)
  r.Put("/{id}", handlers.Update)
  r.Delete("/{id}", handlers.Delete)
  r.Get("/", handlers.List)
  r.Get("/{id}", handlers.Get)

  http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
