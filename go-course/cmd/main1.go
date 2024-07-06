package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nsgowebjavaprog/go-course/pkg/config"
	"github.com/nsgowebjavaprog/go-course/pkg/handlers"
)

const portNumber = ":8083"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot Create template Cache")
	}

	app.TemplateCache = tc

	app.UseCache = false
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Start application on Port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}

// course> go run .
// Start application on Port :8083
