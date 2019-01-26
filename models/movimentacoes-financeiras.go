package models

//MovimentacoesFinanceiras representa as movimentações financeiras de uma determinada pessoa
type MovimentacoesFinanceiras struct {
	ID                  string         `json:"id,omitempty"`
	Pessoa              *Pessoa        `json:"pessoa,omitempty"`
	Valor               float64        `json:"valor,omitempty"`
	Data                string         `json:"data,omitempty"`
	Tipo                string         `json:"tipo,omitempty"`
	Conta               *ContaBancaria `json:"conta,omitempty"`
	Nomeestabelecimento string         `json:"nomeestabelecimento,omitempty"`
}
