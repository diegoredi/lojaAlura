package models

import (
	"lojaAlura/db"
	"strconv"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {

	db := db.ConectaComBancoDeDados()

	selectAll, err := db.Query("Select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectAll.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAll.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func BuscaUmProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()

	p := Produto{}
	var nome, descricao string
	var preco float64
	var quantidade int

	selectOne, err := db.Query("SELECT * FROM produto WHERE id = ", id)

	if err != nil {
		panic(err.Error())
	}

	selectOne.Scan(&id, &nome, &descricao, &preco, &quantidade)

	p.Id, err = strconv.Atoi(id)
	p.Nome = nome
	p.Descricao = descricao
	p.Preco = preco
	p.Quantidade = quantidade

	defer db.Close()

	return p

}

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDados, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insereDados.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletarProduto(id int) {
	db := db.ConectaComBancoDeDados()

	deletaProduto, err := db.Prepare("DELETE FROM produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	deletaProduto.Exec(id)

	defer db.Close()
}

func AtualizarProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	atualizaDados, err := db.Prepare("UPDATE produtos SET nome=$1 , descricao=$2, preco=$3, quantidade=$4 where id = $5")

	if err != nil {
		panic(err.Error())
	}

	atualizaDados.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
