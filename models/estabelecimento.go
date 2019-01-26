package models

//Estabelecimento representa as informações de determinado estabelecimento comercial
type Estabelecimento struct {
	ID       string    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Cnpj     string    `json:"cnpj,omitempty"`
	Ie       string    `json:"ie,omitempty"`
	Endereco *Endereco `json:"endereco,omitempty"`
}
