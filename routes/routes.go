package routes

import (
	"net/http"

	"github.com/jaquelineabreu/pet-loja/controllers"
)

func CarregaRotas(){
	http.HandleFunc("/",controllers.Index)
}