package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB	{
	conexao := "user=postgres dbname=pet_loja password= host=localhost sslmode=disable"
	db, err  := sql.Open("postgres", conexao)
	if err != nil{
		panic(err.Error())
	}

	return db
}

type Produto struct{
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main(){
	db := conectaComBancoDeDados()
	defer db.Close()
	http.HandleFunc("/",index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request){

	produtos := []Produto{
		{"Garrafa de água", "Garrafa de água para cães para passear", 72.00, 1},
		{"Caixa de areia", "Caixa de areia para gatos", 20.00, 2},
		{"Mordedor", "Brinquedo Para Pet Cachorro Pneu Mordedor Resistente Gigante", 30.00, 1},
		{"Wination","Capa de chuva para cachorro - G", 30.00, 2},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}