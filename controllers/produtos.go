package controllers

import (
	"log"
	"lojaAlura/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.BuscaTodosProdutos())
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		urlId := r.URL.Query().Get("id")
		id, err := strconv.Atoi(urlId)
		if err != nil {
			panic(err.Error())
		}

		models.DeletarProduto(id)
		http.Redirect(w, r, "/", 301)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	produto := models.BuscaUmProduto(id)

	temp.ExecuteTemplate(w, "Update", models.AtualizarProduto(produto.Id, produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade))

	http.Redirect(w, r, "/", 301)
}
