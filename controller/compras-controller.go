package controller

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/augustoscher/API-C/model"
	"github.com/gorilla/mux"
)

var compras []model.MovimentacoesCartaoCredito

//AllMovimentoCartaoEndPoint retorna todos os movimentos de cartao
func AllMovimentoCartaoEndPoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(compras)
}

//CreateMovimentoCartaoEndPoint cria um movimento de cartão
func CreateMovimentoCartaoEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var compra model.MovimentacoesCartaoCredito
	_ = json.NewDecoder(r.Body).Decode(&compra)
	compra.ID = params["id"]
	compras = append(compras, compra)
	json.NewEncoder(w).Encode(compras)
}

//FindMovimentoCartaoPorCpfEndpoint retorna movimentos de cartao de credito de determinado cpf
func FindMovimentoCartaoPorCpfEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var movsCartao []model.MovimentacoesCartaoCredito

	for _, item := range compras {
		if item.Pessoa.Cpf == params["cpf"] {
			movsCartao = append(movsCartao, item)
		}
	}
	json.NewEncoder(w).Encode(movsCartao)
}

//AdicionarMovimentosCartaoCredito adiciona alguns movimentos de cartão para servirem de exemplo
func AdicionarMovimentosCartaoCredito() {
	cartao := model.CartaoCredito{
		ID:            "1",
		Nome:          "Nubank",
		Numero:        "79876546898784",
		Dv:            "8",
		Emissao:       "07/2017",
		Validade:      "08/2020",
		Diavencimento: "10",
	}
	cartao2 := model.CartaoCredito{
		ID:            "2",
		Nome:          "Digio",
		Numero:        "8984132494",
		Dv:            "1",
		Emissao:       "04/2015",
		Validade:      "05/2019",
		Diavencimento: "5",
	}
	endereco := model.Endereco{
		ID:     "1",
		Rua:    "Wilhelm Budag",
		Numero: "45",
		Cep:    "89045090",
		Cidade: "Blumenau",
		Estado: "SC",
		Pais:   "Brasil",
	}
	estabelecimento := model.Estabelecimento{
		ID:       "1",
		Nome:     "Meu Estabelecimento",
		Cnpj:     "21555582000142",
		Ie:       "687516546",
		Endereco: &endereco,
	}
	estabelecimento2 := model.Estabelecimento{
		ID:       "2",
		Nome:     "Estabelecimento2",
		Cnpj:     "21123482000142",
		Ie:       "79798798",
		Endereco: &endereco,
	}
	pessoa := pessoas[0]
	pessoa2 := pessoas[1]
	compras = append(compras, model.MovimentacoesCartaoCredito{ID: "1", Pessoa: &pessoa, Valor: 500.00, Data: "26/01/2019", Estabelecimento: &estabelecimento, CartaoCredito: &cartao})
	compras = append(compras, model.MovimentacoesCartaoCredito{ID: "2", Pessoa: &pessoa, Valor: 125.00, Data: "26/01/2019", Estabelecimento: &estabelecimento, CartaoCredito: &cartao})
	compras = append(compras, model.MovimentacoesCartaoCredito{ID: "3", Pessoa: &pessoa2, Valor: 100.00, Data: "26/01/2019", Estabelecimento: &estabelecimento2, CartaoCredito: &cartao2})
}
