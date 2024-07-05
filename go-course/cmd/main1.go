package main

import (
	"fmt"
	"net/http"

	"github.com/nsgowebjavaprog/go-course/pkg/handlers"
)

var portNumber = ":8083"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Start application on Port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}

// course> go run .
// Start application on Port :8083
