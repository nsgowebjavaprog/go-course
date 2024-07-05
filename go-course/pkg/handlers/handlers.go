package handlers

import (
	"net/http"

	"github.com/nsgowebjavaprog/go-course/pkg/render"
)

// home handdler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home_page.html")
}

// About page Handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about_page.html")
}
