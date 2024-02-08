package main

import (
	"fmt"
	"groupie-tracker/internal/routes"
	"net/http"
)

func main() {
	routes.SetupRoutes()

	port := ":8080"

	fmt.Printf("Starting server on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
