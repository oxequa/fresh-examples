package main

import (
	"github.com/tockins/fresh"
	"net/http"
	"github.com/tockins/fresh-examples/models"
)

func main() {
	f := fresh.New()
	f.Config().SetPort(8080)

	// API definition with path and related controller
	f.GET("/todos/", func(c fresh.Context) error{
		todos := [] *models.Todo{}
		todos = append(todos, &models.Todo{})
		todos = append(todos, &models.Todo{})
		return c.Response().JSON(http.StatusOK, todos)
	})
	f.GET("/todos/:uuid", func(c fresh.Context) error{
		todo := models.Todo{c.Request().URLParam("uuid"), "First Todo", "done"}
		return c.Response().JSON(http.StatusOK, todo)
	})
	//Start Fresh Server
	f.Run()
}