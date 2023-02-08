package controllers

import (
	"html/template"
	"net/http"

	"github.com/jaquelineabreu/pet-loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request){
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
	
}

func New(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "New", nil)
}