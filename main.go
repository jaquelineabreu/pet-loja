package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB	{
	conexao := "user=postgres dbname=pet_loja password=minhasenha host=localhost sslmode=disable"
	db, err  := sql.Open("postgres", conexao)
	if err != nil{
		panic(err.Error())
	}

	return db
}

type Produto struct{
	Id int
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main(){	
	http.HandleFunc("/",index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request){
	db := conectaComBancoDeDados()

	selectDetodosOsProdutos, err := db.Query("select * from produtos")

	if err != nil{
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDetodosOsProdutos.Next(){
		var id,quantidade int
		var nome, descricao string
		var preco float64

		err =  selectDetodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil{
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	
	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}