package main

import (
	"net/http"
	"text/template"

	"github.com/jaquelineabreu/pet-loja/models"
)


var temp = template.Must(template.ParseGlob("templates/*.html"))

func main(){	
	http.HandleFunc("/",index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request){
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
	
}