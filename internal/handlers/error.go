package handlers

import (
	"fmt"
	"groupie-tracker/structures"
	"net/http"
	"strconv"
	"text/template"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int, msg string) {
	t, err := template.ParseFiles("ui/templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, strconv.Itoa(http.StatusInternalServerError)+" "+http.StatusText(http.StatusInternalServerError))
		return
	}
	errors := structures.Error{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
	err = t.Execute(w, errors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, strconv.Itoa(http.StatusInternalServerError)+" "+http.StatusText(http.StatusInternalServerError))
		return
	}
}
