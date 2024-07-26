package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Name string
}

func main() {
	tmpl := template.Must(template.New("").ParseGlob("./templates/*"))

	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", PageData{
			Name: "YouTube!",
		})
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Starting website at localhost:8080")

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("An error occured:", err)
	}
}
