package routes

import (
	"groupie-tracker/internal/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	
	fs := http.FileServer(http.Dir("./ui/static"))
	http.Handle("/style.css", fs)
}
