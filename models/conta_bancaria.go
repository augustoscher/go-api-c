package models

//ContaBancaria representa as informações de uma conta bancária de uma pessoa
type ContaBancaria struct {
	ID          string `json:"id,omitempty"`
	Instituicao string `json:"instituicao,omitempty"`
	Agencia     string `json:"agencia,omitempty"`
	Conta       string `json:"conta,omitempty"`
	Digito      string `json:"digito,omitempty"`
}
