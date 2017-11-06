package main

import (
	"github.com/tockins/fresh"
	"github.com/tockins/fresh-examples/models"
	"net/http"
)

func main() {
	f := fresh.New()

	// Main configurations
	f.Config().Port(8080)
	// Setting Default index file
	f.Config().StaticDefault([]string{"index.html"})

	// API definition with path and related controller
	f.GET("/todos/", func(c fresh.Context) error {
		todos := []*models.Todo{}
		todos = append(todos, &models.Todo{})
		todos = append(todos, &models.Todo{})
		return c.Response().JSON(http.StatusOK, todos)
	})
	f.GET("/todos/:uuid", func(c fresh.Context) error {
		todo := models.Todo{c.Request().URLParam("uuid"), "First Todo", "done"}
		return c.Response().JSON(http.StatusOK, todo)
	})

	// Define static assets resources paths
	staticPath := map[string]string{
		"static": "public",
	}
	f.STATIC(staticPath)

	//Start Fresh Server
	f.Run()
}
