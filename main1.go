package main

import (
	"errors"
	"fmt"
	"net/http"
)

var portNumber = ":8084"

// home handdler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey NS-LONI This is Home Page")
}

// About page Handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(10, 10)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This SUM of x & y [10+10]:%d", sum))
}

func AddValues(x, y int) int {
	return x + y
}

// Diivvviiide
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot Divide By 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by Zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Start application on Port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
