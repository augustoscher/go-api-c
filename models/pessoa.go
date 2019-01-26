package models

//Pessoa representa as informações de uma pessoa
type Pessoa struct {
	ID             string    `json:"id,omitempty"`
	Cpf            string    `json:"cpf,omitempty"`
	Datanascimento string    `json:"datanascimento,omitempty"`
	Nome           string    `json:"nome,omitempty"`
	Endereco       *Endereco `json:"endereco,omitempty"`
}
